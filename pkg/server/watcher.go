package server

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Watcher struct {
	target  string
	cli     *clientv3.Client
	builder *ResolverBuilder
}

func NewWatcher(target string, builder *ResolverBuilder) (*Watcher, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(target, ","),
	})
	if err != nil {
		return nil, err
	}

	return &Watcher{cli: cli, builder: builder}, nil
}

func (r *Watcher) Watch() {
	addrDict := make(map[string]map[string]resolver.Address)

	update := func() {
		for k, v := range addrDict {
			addrList := make([]resolver.Address, 0, len(v))
			for _, address := range v {
				fmt.Println("address", address.Addr)
				addrList = append(addrList, address)
			}
			fmt.Println("update", addrList, len(v))
			r.builder.hosts[k] = addrList
			fmt.Println("k", k)
			fmt.Println("r.builder.resolvers[k]", r.builder.resolvers[k])
			fmt.Println("r.builder.resolvers[k]", r.builder.resolvers[k].cc)
			r.builder.resolvers[k].cc.NewAddress(addrList)
		}
	}

	fmt.Println("call etcd", r.target)
	resp, err := r.cli.Get(context.Background(), "/"+Prefix, clientv3.WithPrefix())
	fmt.Println(resp, err)
	fmt.Println(resp.Kvs)
	if err == nil {
		for _, v := range resp.Kvs {
			key := string(v.Key)
			fmt.Println("key", key)
			arr := strings.Split(key, "/")
			fmt.Println(arr, len(arr))
			addresses, ok := addrDict[arr[2]]
			if !ok {
				addresses = make(map[string]resolver.Address)
				addrDict[arr[2]] = addresses
			}
			addresses[string(v.Value)] = resolver.Address{Addr: string(v.Value)}
		}
	}

	update()

	rch := r.cli.Watch(context.Background(), "/"+Prefix, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for n := range rch {
		fmt.Println(n)
		for _, ev := range n.Events {
			fmt.Println("ev", ev.Kv.Key)
			switch ev.Type {
			case mvccpb.PUT:
				key := string(ev.Kv.Key)
				arr := strings.Split(key, "/")
				fmt.Println(arr, len(arr))
				addresses, ok := addrDict[arr[2]]
				if !ok {
					addresses = make(map[string]resolver.Address)
					addrDict[arr[2]] = addresses
				}
				addresses[string(ev.Kv.Value)] = resolver.Address{Addr: string(ev.Kv.Value)}
			case mvccpb.DELETE:
				key := string(ev.PrevKv.Key)
				arr := strings.Split(key, "/")
				fmt.Println(arr, len(arr))
				addresses := addrDict[arr[2]]
				delete(addresses, arr[2])
			}
		}
		update()
	}
}
