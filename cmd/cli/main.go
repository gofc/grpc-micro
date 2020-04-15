package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofc/grpc-micro/pkg/scode"
	"github.com/gofc/grpc-micro/pkg/server"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"github.com/gofc/grpc-micro/proto/v1/pbcomm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"strconv"
	"time"
)

var (
	registryAddress = flag.String("registry_address", "localhost:2379", "registry address")
)

func main() {
	flag.Parse()

	r := server.NewResolverBuilder()
	resolver.Register(r)
	fmt.Println("registry_address", *registryAddress)
	watcher, err := server.NewWatcher(*registryAddress, r)
	if err != nil {
		panic(err)
	}
	go watcher.Watch()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	conn, err := grpc.DialContext(ctx, r.Scheme()+"://authority/"+scode.FOO.Name(),
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name))
	cancel()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(1000 * time.Millisecond)
	for c := range ticker.C {
		fmt.Println("asdfasdf1")

		client := pb.NewFooServiceClient(conn)
		fmt.Println("asdfasdf2")
		res, err := client.Hello(context.Background(), &pbcomm.HelloRequest{
			Name: "gofc" + strconv.Itoa(c.Second()),
		})
		fmt.Println("asdfasdf3")
		fmt.Println(err)
		fmt.Println(res)
	}
}
