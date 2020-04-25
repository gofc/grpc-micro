package services

import (
	"context"
	"fmt"
	"github.com/gofc/grpc-micro/pkg/logger"
	"github.com/gofc/grpc-micro/pkg/server"
	"github.com/gofc/grpc-micro/proto/v1/pbcomm"
)

type FooService struct {
	serverConfig *server.Config
}

func NewFooService(serverConfig *server.Config) *FooService {
	return &FooService{serverConfig: serverConfig}
}

func (f *FooService) Hello(ctx context.Context, req *pbcomm.HelloRequest) (*pbcomm.HelloResponse, error) {
	logger.Debugf("address %s, hello %s", f.serverConfig.Address, req.Name)
	return &pbcomm.HelloResponse{
		Msg: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}
