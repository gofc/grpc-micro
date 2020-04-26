package user

import "github.com/gofc/grpc-micro/internal/user/services"

//Server of api services
type Server struct {
	UserService *services.UserService
}

func NewUserServer() *Server {
	return &Server{
		services.NewUserService(),
	}
}
