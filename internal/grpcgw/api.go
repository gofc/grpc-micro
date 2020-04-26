package grpcgw

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"github.com/gofc/grpc-micro/pkg/registry/etcd"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

//Server of api services
type Server struct {
	userServiceClient pb.UserServiceClient
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return s.userServiceClient.Login(ctx, req)
}

func (s *Server) Me(ctx context.Context, req *pb.MeRequest) (*pb.MeResponse, error) {
	return s.userServiceClient.Me(ctx, req)
}

func NewGatewayServer() (*Server, error) {
	//todo registry resolver
	cli, err := clientv3.NewFromURL(Conf.Registry.Address)
	if err != nil {
		return nil, err
	}
	r := &etcdnaming.GRPCResolver{Client: cli}

	resolverBuilder := etcd.NewResolverBuilder()

	resolver.Register(resolverBuilder)

	resolver.SetDefaultScheme("")

	conn, err := grpc.Dial("my-service",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	userServiceClient := pb.NewUserServiceClient(conn)

	return &Server{
		userServiceClient: userServiceClient,
	}, nil
}
