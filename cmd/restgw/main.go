package main

import (
	"context"
	"flag"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var (
	httpAddr = flag.String("port", ":40000", "http server listen addr")
	endpoint = flag.String("endpoint", "localhost:50000", "grpc gateway address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{OrigName: false, EmitDefaults: true},
		),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}

	_ = pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *endpoint, opts)

	return http.ListenAndServe(*httpAddr, mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
