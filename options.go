package server

import (
	"net/http"

	"google.golang.org/grpc"
)

//Option type
type Option func(*Options)

//Options of Server
type Options struct {
	Name         string
	Address      string
	Port         int
	Kuberesolver bool

	LoggingDecider         func(fullMethodName string) bool
	AdditionalInterceptors []grpc.UnaryServerInterceptor

	HTTPPort    int
	HTTPHandler http.Handler

	BeforeStop  func() error
	GRPCHandler RegisterGRPCService
}

//WithServiceName set service name
func WithServiceName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

//WithBeforeStop set a callback before stop Server
func WithBeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = fn
	}
}

//WithAddress set gRPC Server listening address
func WithAddress(address string) Option {
	return func(o *Options) {
		o.Address = address
	}
}

//WithLoggingDecider set logger filter
func WithLoggingDecider(pd func(fullMethodName string) bool) Option {
	return func(o *Options) {
		o.LoggingDecider = pd
	}
}

//WithAdditionalInterceptors set additional interceptors
func WithAdditionalInterceptors(args ...grpc.UnaryServerInterceptor) Option {
	return func(o *Options) {
		o.AdditionalInterceptors = args
	}
}

//WithHTTPHandler set http handler for http/1.1 access
func WithHTTPHandler(port int, handler http.Handler) Option {
	return func(o *Options) {
		o.HTTPPort = port
		o.HTTPHandler = handler
	}
}

//WithPort set gRPC listening port
func WithPort(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}

//WithKuberesolver turn on/off service discovery within k8s
func WithKuberesolver(enable bool) Option {
	return func(o *Options) {
		o.Kuberesolver = enable
	}
}

//WithRegisterGRPCService set GRPCHandler
func WithRegisterGRPCService(fn RegisterGRPCService) Option {
	return func(o *Options) {
		o.GRPCHandler = fn
	}
}
