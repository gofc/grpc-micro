package code

import (
	"context"
	"fmt"
	"github.com/gofc/grpc-micro/logger"
	"sync"

	"go.uber.org/zap"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

//ServiceCode an unique service code
type ServiceCode string

//Config grpc connection target
type Config struct {
	Target string
}

//Name get service code
func (s ServiceCode) Name() string {
	return string(s)
}

//ConnPool コネクションプール
type ConnPool struct {
	lock    sync.Mutex
	conns   map[ServiceCode]*grpc.ClientConn
	targets map[string]*Config
}

//Get get service grpc connection
func (cp *ConnPool) Get(code ServiceCode) *grpc.ClientConn {
	cp.lock.Lock()
	defer cp.lock.Unlock()
	if conn, ok := cp.conns[code]; ok {
		return conn
	}

	c, ok := cp.targets[code.Name()]
	if !ok {
		panic(fmt.Sprintf("%s service target not found", code.Name()))
	}
	tracer := opentracing.GlobalTracer()
	conn, err := grpc.Dial(c.Target,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithChainUnaryInterceptor(
			func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
				md := metautils.ExtractIncoming(ctx).Clone()
				return invoker(md.ToOutgoing(ctx), method, req, reply, cc, opts...)
			},
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(tracer)),
		),
		grpc.WithChainStreamInterceptor(
			func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer,
				opts ...grpc.CallOption) (grpc.ClientStream, error) {
				md := metautils.ExtractIncoming(ctx).Clone()
				return streamer(md.ToOutgoing(ctx), desc, cc, method, opts...)
			},
			grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(tracer)),
		),
	)
	if err != nil {
		panic(err)
	}
	cp.conns[code] = conn
	return conn
}

//Close コネクションを開放する
func (cp *ConnPool) Close(ctx context.Context) {
	for code, conn := range cp.conns {
		logger.Info(ctx, "release grpc connection", zap.String("ServiceCode", code.Name()))
		if err := conn.Close(); err != nil {
			logger.Error(ctx, "failed to close auth connection", zap.Error(err))
		}
	}
}

//NewConnPool 初期化
func NewConnPool(targets map[string]*Config) *ConnPool {
	return &ConnPool{
		lock:    sync.Mutex{},
		conns:   make(map[ServiceCode]*grpc.ClientConn),
		targets: targets,
	}
}
