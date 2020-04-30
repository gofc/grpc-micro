package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/gofc/grpc-micro/pkg/logger"
	"github.com/gofc/grpc-micro/pkg/registry"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
	"sync"
)

type etcdResolverBuilder struct {
	Client *clientv3.Client
	r      *etcdResolver
}

func (b *etcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	if b.r == nil {
		ctx, cancelFunc := context.WithCancel(context.Background())
		r := &etcdResolver{
			ctx:    ctx,
			cancel: cancelFunc,
			lock:   sync.RWMutex{},
			target: target.Endpoint,
			c:      b.Client,
		}
		r.ccs = make(map[string]resolver.ClientConn)
		r.svcs = make(map[string]*registry.Service)
		b.r = r

		go r.watcher()
	}
	b.r.ccs[target.Endpoint] = cc
	return b.r, nil
}

func (b *etcdResolverBuilder) Scheme() string {
	return schema
}

func NewResolverBuilder(registryAddress string) (*etcdResolverBuilder, error) {
	cli, err := clientv3.NewFromURL(registryAddress)
	if err != nil {
		return nil, err
	}
	return &etcdResolverBuilder{Client: cli}, nil
}

type etcdResolver struct {
	ctx    context.Context
	cancel context.CancelFunc
	lock   sync.RWMutex

	target string

	c    *clientv3.Client
	svcs map[string]*registry.Service
	ccs  map[string]resolver.ClientConn
}

func (e *etcdResolver) appendTarget(target resolver.Target, cc resolver.ClientConn) {
	e.ccs[target.Endpoint] = cc
}

func (e *etcdResolver) ResolveNow(resolver.ResolveNowOptions) {
	res, err := e.c.Get(e.ctx, e.target, clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		logger.Error("etcd: failed get targets", zap.Error(err))
		return
	}

	e.lock.Lock()
	defer e.lock.Unlock()

	e.svcs = make(map[string]*registry.Service)

	for _, kv := range res.Kvs {
		sn := decode(kv.Value)
		if sn == nil {
			continue
		}
		e.svcs[sn.ID] = sn
	}
	e.flushClientConn()
}

func (e *etcdResolver) Close() {
	e.cancel()
}

func (e *etcdResolver) watcher() {
	//todo watcher实现
	logger.CInfo(context.Background(), "start etcd registry watcher")
}

func (e *etcdResolver) flushClientConn() {
	m := make(map[string][]resolver.Address)
	for _, sn := range e.svcs {
		addrs, ok := m[sn.Name]
		if !ok {
			addrs = make([]resolver.Address, 0, 0)
		}
		addrs = append(addrs, resolver.Address{
			Addr: sn.Address,
		})
		m[sn.Name] = addrs
	}
	for name, addrs := range m {
		cc, ok := e.ccs[name]
		if !ok {
			continue
		}
		cc.UpdateState(resolver.State{
			Addresses: addrs,
			//todo ServiceConfig测试
		})
	}
}
