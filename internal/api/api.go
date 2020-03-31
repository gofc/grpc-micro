package api

import "github.com/gofc/grpc-micro/internal/api/services"

type APIServer struct {
	FooService *services.FooService
}

func NewAPIServer() *APIServer {
	return &APIServer{
		services.NewFooService(),
	}
}
