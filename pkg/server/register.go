package server

import (
	"github.com/gofc/grpc-micro/pkg/scode"
)

// Prefix should start and end with no slash
var Prefix = "etcd3_naming"

//var client clientv3.Client
//var serviceKey string
//var stopSignal = make(chan bool, 1)

// Register
func Register(serviceCode scode.ServiceCode, address string, registryAddress string) error {
	//name := serviceCode.Name()
	//var err error
	//var host, port string
	//
	//if cnt := strings.Count(address, ":"); cnt >= 1 {
	//	// ipv6 address in format [host]:port or ipv4 host:port
	//	host, port, err = net.SplitHostPort(address)
	//	if err != nil {
	//		return err
	//	}
	//} else {
	//	host = address
	//}
	//ipAddr, err := addr.Extract(host)
	//if err != nil {
	//	return err
	//}
	//
	//serviceValue := fmt.Sprintf("%s:%s", ipAddr, port)
	//serviceKey = fmt.Sprintf("/%s/%s/%s", Prefix, name, serviceValue)
	//
	//logger.Debugf("register server, key:%v, value:%v", serviceKey, serviceValue)
	//
	//// get endpoints for register dial address
	//client, err := clientv3.New(clientv3.Config{
	//	Endpoints: strings.Split(registryAddress, ","),
	//})
	//if err != nil {
	//	return errors.Wrap(err, "create etcd3 client failed")
	//}
	//go func() {
	//	// invoke self-register with ticker
	//	ticker := time.NewTicker(time.Second * 10)
	//	for {
	//		// minimum lease TTL is ttl-second
	//		resp, _ := client.Grant(context.TODO(), 15)
	//		// should get first, if not exist, set it
	//		_, err := client.Get(context.Background(), serviceKey)
	//		if err != nil {
	//			if err == rpctypes.ErrKeyNotFound {
	//				if _, err := client.Put(context.TODO(), serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
	//					logger.Error("failed to set service ttl", zap.String("name", name), zap.Error(err))
	//				}
	//			} else {
	//				logger.Error("failed to connect registry server", zap.String("name", name), zap.Error(err))
	//			}
	//		} else {
	//			// refresh set to true for not notifying the watcher
	//			if _, err := client.Put(context.Background(), serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
	//				logger.Error("failed to refresh service ttl", zap.String("name", name), zap.Error(err))
	//			}
	//		}
	//		select {
	//		case <-stopSignal:
	//			return
	//		case <-ticker.C:
	//		}
	//	}
	//}()
	return nil
}

// UnRegister delete registered service from etcd
func UnRegister() error {
	//stopSignal <- true
	//stopSignal = make(chan bool, 1) // just a hack to avoid multi UnRegister deadlock
	//var err error
	//if _, err := client.Delete(context.Background(), serviceKey); err != nil {
	//	logger.Error("failed to deregister", zap.String("serviceKey", serviceKey), zap.Error(err))
	//} else {
	//	logger.CInfo(context.Background(), "deregister finished", zap.String("serviceKey", serviceKey))
	//}
	//return err
	return nil
}
