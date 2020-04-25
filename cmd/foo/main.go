package main

import (
	"context"
	"flag"
	"github.com/gofc/grpc-micro/internal/foo"
	"github.com/gofc/grpc-micro/pkg/logger"
	"github.com/gofc/grpc-micro/pkg/scode"
	"github.com/gofc/grpc-micro/pkg/server"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	envName    = flag.String("env", "local", "environment name")
	configPath = flag.String("conf", "", "the folder path of config files")
)

func init() {
	flag.Parse()

	code := scode.FOO
	var filePath string
	if *configPath == "" {
		filePath = filepath.Clean(
			filepath.Join("../configs/app", *envName, code.Name()+".yml"),
		)
	} else {
		filePath = filepath.Clean(
			filepath.Join(*configPath, *envName, code.Name()+".yml"),
		)
	}
	if err := configor.Load(foo.Conf, filePath); err != nil {
		panic(err)
	}
}

func main() {
	conf := foo.Conf

	logger.InitLogger(conf.Logger)
	ctx := context.Background()
	grpcServer := grpc.NewServer()
	apiServer := foo.NewAPIServer()

	pb.RegisterFooServiceServer(grpcServer, apiServer.FooService)

	ts, err := net.Listen("tcp", conf.Server.Address)
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
	if err = server.Register(scode.FOO, ts.Addr().String(), conf.Registry.Address); err != nil {
		logger.Error("failed to register service", zap.Error(err))
		return
	}

	if err = grpcServer.Serve(ts); err != nil {
		logger.Error("failed to start server", zap.Error(err))
		return
	}
}
