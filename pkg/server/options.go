package server

import (
	"context"
	"github.com/gofc/grpc-micro/pkg/registry"
	"time"

	"google.golang.org/grpc"
)

type Option func(*Options)

type Options struct {
	Registry registry.Registry
	Name     string
	Address  string
	Id       string
	Version  string

	// RegisterCheck runs a check function before registering the service
	RegisterCheck func(context.Context) error
	// The register expiry time
	RegisterTTL time.Duration
	// The interval on which to register
	RegisterInterval time.Duration

	GrpcOptions []grpc.ServerOption
	MaxMsgSize  int
}

func newOptions(opt ...Option) Options {
	opts := Options{
		RegisterInterval: DefaultRegisterInterval,
		RegisterTTL:      DefaultRegisterTTL,
	}

	for _, o := range opt {
		o(&opts)
	}

	if opts.Registry == nil {
		opts.Registry = registry.DefaultRegistry
	}

	if opts.RegisterCheck == nil {
		opts.RegisterCheck = DefaultRegisterCheck
	}

	if len(opts.Address) == 0 {
		opts.Address = DefaultAddress
	}

	if len(opts.Name) == 0 {
		opts.Name = DefaultName
	}

	if len(opts.Id) == 0 {
		opts.Id = DefaultId
	}

	if len(opts.Version) == 0 {
		opts.Version = DefaultVersion
	}

	return opts
}

// Server name
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Unique server id
func Id(id string) Option {
	return func(o *Options) {
		o.Id = id
	}
}

// Version of the service
func Version(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

// Address to bind to - host:port
func Address(a string) Option {
	return func(o *Options) {
		o.Address = a
	}
}

// Registry used for discovery
func Registry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

// RegisterCheck run func before registry service
func RegisterCheck(fn func(context.Context) error) Option {
	return func(o *Options) {
		o.RegisterCheck = fn
	}
}

// Register the service with a TTL
func RegisterTTL(t time.Duration) Option {
	return func(o *Options) {
		o.RegisterTTL = t
	}
}

// Register the service with at interval
func RegisterInterval(t time.Duration) Option {
	return func(o *Options) {
		o.RegisterInterval = t
	}
}
