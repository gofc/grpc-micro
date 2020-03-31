package services

import (
	"context"
	"fmt"
	"github.com/gofc/grpc-micro/pkg/logger"
	"github.com/gofc/grpc-micro/proto/v1/pbcomm"
)

type FooService struct {
}

func NewFooService() *FooService {
	return &FooService{}
}

func (f *FooService) Hello(ctx context.Context, req *pbcomm.HelloRequest) (*pbcomm.HelloResponse, error) {
	logger.Debugf("hello %s", req.Name)
	return &pbcomm.HelloResponse{
		Msg: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}
