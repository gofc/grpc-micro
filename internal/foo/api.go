package foo

import "github.com/gofc/grpc-micro/internal/foo/services"

//Server of api services
type Server struct {
	FooService *services.FooService
}

func NewAPIServer() *Server {
	return &Server{
		services.NewFooService(Conf.Server),
	}
}
