package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc/resolver"
)

const schema = "grpc-micro-registry"

type ResolverBuilder struct {
	Client *clientv3.Client
}

func NewResolverBuilder(registryAddress string) (*ResolverBuilder, error) {
	cli, err := clientv3.NewFromURL(registryAddress)
	if err != nil {
		return nil, err
	}
	return &ResolverBuilder{Client: cli}, nil
}

func (r *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	panic("implement me")
}

func (r *ResolverBuilder) Scheme() string {
	return schema
}
