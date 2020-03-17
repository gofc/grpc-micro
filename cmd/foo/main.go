package main

import (
	"fmt"
	"github.com/gofc/grpc-micro/pkg/logger"
	"github.com/gofc/grpc-micro/pkg/registry"
	"github.com/gofc/grpc-micro/pkg/registry/etcd"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger.InitLogger(&logger.Config{
		Level:  "debug",
		Format: "text",
	})
	registry.DefaultRegistry = etcd.NewRegistry()

	srv := grpc.NewServer()

	l, err := net.Listen("tcp", ":50000")
	if err != nil {
		logger.Error("failed to listen port", zap.Error(err))
		return
	}

	if err = srv.Serve(l); err != nil {
		logger.Error("failed to start server", zap.Error(err))
		return
	}

	fmt.Println("HelloWorld!!")
}
