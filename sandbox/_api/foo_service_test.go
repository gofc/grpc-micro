package _api

import (
	"context"
	"fmt"
	pb "github.com/gofc/grpc-micro/proto/v1"
	"github.com/gofc/grpc-micro/proto/v1/pbcomm"
	"google.golang.org/grpc"
	"testing"
)

func TestFooService_Hello(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "localhost:50001", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewFooServiceClient(conn)

	res, err := client.Hello(context.Background(), &pbcomm.HelloRequest{
		Name: "gofc",
	})
	fmt.Println(err)
	fmt.Println(res)
}
