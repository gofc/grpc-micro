package main

import (
	"context"
	"github.com/gofc/grpc-micro/internal/api"
	"github.com/gofc/grpc-micro/pkg/logger"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger.InitLogger(&logger.Config{
		Level:  "debug",
		Format: "text",
	})
	ctx := context.Background()
	grpcServer := grpc.NewServer()
	apiServer := api.NewAPIServer()

	pb.RegisterFooServiceServer(grpcServer, apiServer.FooService)

	l, err := net.Listen("tcp", ":50000")
	if err != nil {
		logger.Error("failed to listen port", zap.Error(err))
		return
	}
	logger.CInfo(ctx, "server start to listening", zap.String("address", ":50000"))
	if err = grpcServer.Serve(l); err != nil {
		logger.Error("failed to start server", zap.Error(err))
		return
	}
}
