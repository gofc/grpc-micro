package micro

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gofc/grpc-micro/kuberesolver"
	"github.com/gofc/grpc-micro/logger"

	"github.com/gofc/grpc-micro/code"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//RegisterGRPCService register grpc services
type RegisterGRPCService interface {
	Register(ctx context.Context, srv *grpc.Server) error
}

//Server RPC server
type Server struct {
	options Options
	srv     *grpc.Server
}

//Server returns gRPC Server
func (s *Server) Server() *grpc.Server {
	return s.srv
}

//Start Server
func (s *Server) Start(ctx context.Context, opts ...Option) error {
	for _, o := range opts {
		o(&s.options)
	}

	if s.options.GRPCHandler == nil {
		return errors.New("failed to get grpc handler settings")
	}
	if err := s.options.GRPCHandler.Register(ctx, s.srv); err != nil {
		return err
	}

	var hts net.Listener
	ts, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.options.Address, s.options.Port))
	if err != nil {
		return err
	}
	if s.options.HTTPHandler != nil {
		hts, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.options.Address, s.options.HTTPPort))
		if err != nil {
			return err
		}
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		sig := <-ch
		logger.Info(ctx, "receive signal", zap.String("signal", sig.String()))
		if err := s.options.BeforeStop(); err != nil {
			logger.Error(ctx, "failed call BeforeStop method", zap.Error(err))
		}
		s.srv.GracefulStop()
		os.Exit(1)
	}()

	if hts != nil {
		go func() {
			logger.Info(ctx, "http Server start listening",
				zap.String("address", s.options.Address),
				zap.Int("port", s.options.HTTPPort))
			if err := http.Serve(hts, customMimeWrapper(s.options.HTTPHandler)); err != nil {
				panic(err)
			}
		}()
	}

	logger.Info(ctx, "grpc Server start listening",
		zap.String("address", s.options.Address),
		zap.Int("port", s.options.Port))
	if err := s.srv.Serve(ts); err != nil {
		logger.Error(ctx, "failed start grpc Server", zap.Error(err))
		panic(err)
	}
	return nil
}

func customMimeWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "webhook") {
			r.Header.Set("Content-Type", "application/raw-webhook")
		}
		h.ServeHTTP(w, r)
	})
}

//NewServer create instance
func NewServer(opts ...Option) *Server {
	s := &Server{
		options: Options{
			Address:  "0.0.0.0",
			Port:     50000,
			HTTPPort: 31000,
			LoggingDecider: func(fullMethodName string) bool {
				return fullMethodName != "/grpc.health.v1.Health/Check"
			},
			BeforeStop: func() error { return nil },
		},
	}

	for _, o := range opts {
		o(&s.options)
	}

	interceptors := make([]grpc.UnaryServerInterceptor, 0)

	interceptors = append(interceptors,
		grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithFilterFunc(func(ctx context.Context, fullMethodName string) bool {
			//ignore health check
			return fullMethodName != "/grpc.health.v1.Health/Check"
		})),
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
			sp := opentracing.SpanFromContext(ctx)
			if sp != nil {
				if spanContext, ok := sp.Context().(ddtrace.SpanContext); ok {
					ctx = grpc_ctxtags.SetInContext(ctx, grpc_ctxtags.NewTags().
						Set("dd.trace_id", spanContext.TraceID()).
						Set("dd.span_id", spanContext.SpanID()))
				}
			}
			return handler(ctx, req)
		},
		grpc_zap.PayloadUnaryServerInterceptor(
			logger.GetLogger(),
			func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
				return s.options.LoggingDecider(fullMethodName)
			}),
		grpc_zap.UnaryServerInterceptor(logger.GetLogger(), grpc_zap.WithDecider(func(fullMethodName string, err error) bool {
			return s.options.LoggingDecider(fullMethodName)
		})),
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					logger.Error(ctx, "panic error", zap.Error(err))
					err = status.Error(codes.Internal, "system error")
				}
			}()
			return handler(ctx, req)
		},
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			resp, err = handler(ctx, req)
			if err == nil {
				return
			}

			if v, ok := err.(code.ErrorCode); ok {
				err = v.GrpcError(codes.FailedPrecondition)
				logger.ErrorNoStack(ctx, "ecode error", zap.String("error", v.Error()))
				return
			}

			switch v := err.(type) {
			case *code.AppError:
				err = code.NewGRPC(codes.FailedPrecondition, err)
				logger.ErrorNoStack(ctx, "application error", zap.String("error", v.Error()))
			case *code.GrpcError:
				logger.ErrorNoStack(ctx, "ecode grpc error", zap.String("error", v.Error()))
				return
			default:
				logger.ErrorNoStack(ctx, "grpc error", zap.Error(v))
				st, ok := status.FromError(err)
				if ok {
					return nil, st.Err()
				}
				err = code.NewGRPC(codes.Internal, err)
			}
			return
		},
	)

	if len(s.options.AdditionalInterceptors) > 0 {
		interceptors = append(interceptors, s.options.AdditionalInterceptors...)
	}

	s.srv = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)),
	)

	//add health check
	hsrv := health.NewServer()
	hsrv.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s.srv, hsrv)

	//add Kuberesolver
	if s.options.Kuberesolver {
		logger.Info(context.Background(), "start k8s registry...")
		kuberesolver.RegisterInCluster()
	}

	return s
}
