package _api

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/gofc/grpc-micro/pkg/scode"
	"github.com/gofc/grpc-micro/pkg/server"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"github.com/gofc/grpc-micro/proto/v1/pbcomm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFooService_Hello(t *testing.T) {
	r := server.NewResolverBuilder()
	resolver.Register(r)
	watcher, err := server.NewWatcher("localhost:2379")
	if err != nil {
		panic(err)
	}
	go watcher.Watch()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	conn, err := grpc.DialContext(ctx, "etcd3_naming://authority/"+scode.FOO.Name(),
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

func TestETCD(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split("localhost:2379", ","),
	})
	fmt.Println(err)
	res, err := cli.Get(context.Background(), "/etcd3_naming", clientv3.WithPrefix())
	fmt.Println(res, err)
	fmt.Println(res.Kvs)
}
