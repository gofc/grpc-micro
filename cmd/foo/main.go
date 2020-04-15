package main

import (
	"context"
	"flag"
	"github.com/gofc/grpc-micro/internal/api"
	"github.com/gofc/grpc-micro/pkg/logger"
	"github.com/gofc/grpc-micro/pkg/scode"
	"github.com/gofc/grpc-micro/pkg/server"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	address         = flag.String("address", ":0", "listening host")
	registry        = flag.String("registry", "etcd", "registry type")
	registryAddress = flag.String("registry_address", "localhost:2379", "registry address")
)

func main() {
	flag.Parse()

	logger.InitLogger(&logger.Config{
		Level:  "debug",
		Format: "text",
	})
	ctx := context.Background()
	grpcServer := grpc.NewServer()
	apiServer := api.NewAPIServer()

	pb.RegisterFooServiceServer(grpcServer, apiServer.FooService)

	ts, err := net.Listen("tcp", *address)
	if err != nil {
		logger.Error("failed to listen port", zap.Error(err))
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		logger.Error("receive signal", zap.String("signal", s.String()))
		grpcServer.GracefulStop()
		if err := server.UnRegister(); err != nil {
			logger.Error("unregister failed", zap.Error(err))
		}
		os.Exit(1)
	}()

	logger.CInfo(ctx, "server start to listening", zap.String("address", ts.Addr().String()))
	if err = server.Register(scode.FOO, ts.Addr().String(), *registryAddress); err != nil {
		logger.Error("failed to register service", zap.Error(err))
		return
	}

	if err = grpcServer.Serve(ts); err != nil {
		logger.Error("failed to start server", zap.Error(err))
		return
	}
}
