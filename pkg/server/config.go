package server

//Config of gRPC server
type Config struct {
	Address string `yaml:"address" required:"true"`
}
