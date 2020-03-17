package server

import "google.golang.org/grpc"

type grpcServer struct {
	name    string
	version string

	srv *grpc.Server
}

func NewServer() *grpcServer {
	return &grpcServer{}
}
