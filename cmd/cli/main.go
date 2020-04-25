package main

import (
	"flag"
)

var (
	registryAddress = flag.String("registry_address", "localhost:2379", "registry address")
)

func main() {
	flag.Parse()

	//r := server.NewResolverBuilder()
	//resolver.Register(r)
	//fmt.Println("registry_address", *registryAddress)
	//watcher, err := server.NewWatcher(*registryAddress, r)
	//if err != nil {
	//	panic(err)
	//}
	//go watcher.Watch()

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//
	//conn, err := grpc.DialContext(ctx, r.Scheme()+"://authority/"+scode.FOO.Name(),
	//	grpc.WithInsecure(),
	//	grpc.WithBalancerName(roundrobin.Name))
	//cancel()
	//if err != nil {
	//	panic(err)
	//}
	//client := pb.NewFooServiceClient(conn)
	//
	//ticker := time.NewTicker(1000 * time.Millisecond)
	//for c := range ticker.C {
	//	res, err := client.Hello(context.Background(), &pbcomm.HelloRequest{
	//		Name: "gofc" + strconv.Itoa(c.Second()),
	//	})
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println(res)
	//	}
	//}
}
