package server

import (
	"fmt"
	"google.golang.org/grpc/resolver"
)

type ResolverBuilder struct {
	hosts     map[string][]resolver.Address
	resolvers map[string]*Resolver
}

func NewResolverBuilder() *ResolverBuilder {
	return &ResolverBuilder{
		hosts:     make(map[string][]resolver.Address),
		resolvers: make(map[string]*Resolver),
	}
}

// Scheme return etcdv3 schema
func (r *ResolverBuilder) Scheme() string {
	return Prefix
}

// Build to resolver.Resolver
func (r *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	fmt.Println(target.Endpoint)
	res, ok := r.resolvers[target.Endpoint]
	fmt.Println("ok", ok)
	if !ok {
		res = NewResolver(target.Endpoint)
		r.resolvers[target.Endpoint] = res
	}
	cc.NewAddress(r.hosts[target.Endpoint])
	res.cc = cc

	fmt.Println("r.resolvers", r.resolvers)

	return res, nil
}

// resolver is the implementaion of grpc.resolve.Builder
type Resolver struct {
	service string
	cc      resolver.ClientConn
}

// NewResolver return resolver builder
// target example: "http://127.0.0.1:2379,http://127.0.0.1:12379,http://127.0.0.1:22379"
// service is service name
func NewResolver(service string) *Resolver {
	return &Resolver{service: service}
}

// ResolveNow
func (r *Resolver) ResolveNow(rn resolver.ResolveNowOption) {
}

// Close
func (r *Resolver) Close() {
}
