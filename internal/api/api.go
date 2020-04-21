package api

import "github.com/gofc/grpc-micro/internal/api/services"

//Server of api services
type Server struct {
	FooService *services.FooService
}

func NewAPIServer() *Server {
	return &Server{
		services.NewFooService(),
	}
}
