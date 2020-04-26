package user

import (
	"github.com/gofc/grpc-micro/pkg/logger"
	"github.com/gofc/grpc-micro/pkg/registry"
	"github.com/gofc/grpc-micro/pkg/server"
)

var (
	//Conf of api services
	Conf = &Config{}
)

//Config of system
type Config struct {
	Server   *server.Config
	Logger   *logger.Config
	Registry *registry.Config
}
