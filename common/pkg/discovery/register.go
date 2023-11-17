package discovery

import (
	"context"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type EtcdRegister struct {
	//client config
	EtcdAddr    []string
	DialTimeout time.Duration

	closeCh     chan struct{}
	leaseId     clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse

	servInfo GrpcServer
	servTTL  int64
	client   *clientv3.Client
}

func NewRegister(EtcdAddr []string) *EtcdRegister {
	return &EtcdRegister{
		EtcdAddr:    EtcdAddr,
		DialTimeout: 5 * time.Second,
	}
}

func (er *EtcdRegister) Register(server GrpcServer, ttl int64) (chan<- struct{}, error) {
	// 创建etcd客户端连接
	var err error
	if er.client, err = clientv3.New(clientv3.Config{
		Endpoints:   er.EtcdAddr, // etcd服务器的地址
		DialTimeout: er.DialTimeout,
	}); err != nil {
		return nil, err
	}

	er.servInfo = server
	er.servTTL = ttl

	if err = er.doRegister(); err != nil {
		return nil, err
	}
	er.closeCh = make(chan struct{})
	go er.keepAlive()

	return er.closeCh, nil
}

func (er *EtcdRegister) Stop() {
	er.closeCh <- struct{}{}
}

func (er *EtcdRegister) doRegister() error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), er.DialTimeout*time.Second)
	defer cancel()

	// 创建一个租约，用于定期刷新注册信息
	leaseResp, err := er.client.Grant(ctx, er.servTTL) // 租约时间为10秒
	if err != nil {
		return err
	}
	er.leaseId = leaseResp.ID
	if er.keepAliveCh, err = er.client.KeepAlive(context.Background(), er.leaseId); err != nil {
		return err
	}

	// 服务注册的key和value
	serviceKey := "/" + er.servInfo.Name
	// 你的服务地址
	serviceValue, _ := json.Marshal(er.servInfo)
	// 将服务信息注册到etcd
	_, err = er.client.Put(context.Background(), serviceKey, string(serviceValue), clientv3.WithLease(er.leaseId))
	if err != nil {
		return err
	}
	return nil
}

func (er *EtcdRegister) unregister() error {
	_, err := er.client.Delete(context.Background(), "/"+er.servInfo.Name)
	return err
}

func (er *EtcdRegister) keepAlive() {
	ticker := time.NewTicker(time.Duration(er.servTTL) * time.Second)
	for {
		select {
		case <-er.closeCh:
			if err := er.unregister(); err != nil {
				log.Println(er.servInfo.Name+" unregister failed, error: ", err)
			}

			if _, err := er.client.Revoke(context.Background(), er.leaseId); err != nil {
				log.Println(er.servInfo.Name+" revoke failed, error: ", err)
			}
		case res := <-er.keepAliveCh:
			if res == nil {
				if err := er.doRegister(); err != nil {
					log.Println(er.servInfo.Name + " doRegister fail...")
				}
			}
			ticker.Reset(time.Duration(er.servTTL) * time.Second)
			//ticker = time.NewTicker(time.Duration(er.servTTL) * time.Second)
		case <-ticker.C:
			if er.keepAliveCh != nil {
				if err := er.doRegister(); err != nil {
					log.Println(er.servInfo.Name + " doRegister fail...")
				}
			}

		}
	}
}
