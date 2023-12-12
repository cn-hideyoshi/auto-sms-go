package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	clientV3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

type EtcdResolver struct {
	schema      string
	EtcdAddr    []string
	DialTimeout time.Duration

	target resolver.Target

	closeCh      chan struct{}
	watchCh      clientV3.WatchChan
	client       *clientV3.Client
	keyPrefix    string
	servAddrList []resolver.Address

	cc resolver.ClientConn
}

func NewResolver(EtcdAddr []string) *EtcdResolver {
	return &EtcdResolver{
		schema:      "etcd",
		EtcdAddr:    EtcdAddr,
		DialTimeout: 5 * time.Second,
	}
}

func (er *EtcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// 创建 etcd 客户端连接
	etcdClient, err := clientV3.New(clientV3.Config{
		Endpoints: er.EtcdAddr,
	})
	if err != nil {
		return nil, err
	}
	er.cc = cc
	er.client = etcdClient
	er.target = target
	resolver.Register(er)

	er.closeCh = make(chan struct{})

	//初始化地址
	er.queryEtcd()

	// 启动 Resolver 的更新循环
	go er.watchUpdates()
	return er, nil
}

func (er *EtcdResolver) watchUpdates() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-er.closeCh:
		case <-ticker.C:
			err := er.queryEtcd()
			if err != nil {
				log.Printf("Error querying etcd: %v", err)
				// 在发生错误时等待一段时间，然后重试
				time.Sleep(5 * time.Second)
				continue
			}
			// 更新 ClientConn 中的服务地址
			er.cc.UpdateState(resolver.State{
				Addresses: er.servAddrList,
			})
		}
	}
}

func (er *EtcdResolver) queryEtcd() error {
	key := fmt.Sprintf("/%s", er.target.Endpoint())
	resp, err := er.client.Get(context.Background(), key, clientV3.WithPrefix())
	if err != nil {
		return err
	}

	var addrs []resolver.Address
	for _, kv := range resp.Kvs {
		info := GrpcServer{}
		_ = json.Unmarshal(kv.Value, &info)
		addrs = append(addrs, resolver.Address{
			Addr:       info.Addr,
			ServerName: info.Name,
		})
	}
	er.servAddrList = addrs
	return err
}

func (er *EtcdResolver) ResolveNow(o resolver.ResolveNowOptions) {}

func (er *EtcdResolver) Scheme() string {
	return er.schema
}

func (er *EtcdResolver) Close() {
}
