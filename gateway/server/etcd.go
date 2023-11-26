package server

import (
	"blog.hideyoshi.top/common/pkg/discovery"
	"blog.hideyoshi.top/gateway/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func NewRegister() {
	etcdAddr := []string{config.C.Etcd.Addr}
	register := discovery.NewRegister(etcdAddr)

	server := discovery.GrpcServer{
		Name: config.C.Server.Name,
		Addr: config.C.Server.Addr,
	}

	_, err := register.Register(server, 10)
	if err != nil {
		log.Fatalln("register etcd fail :", err)
	}
}

func RegisterGateway() {
	c := GrpcConfig{
		Addr: config.C.Server.Addr,
		RegisterFunc: func(server *grpc.Server) {
			//no gateway server
		},
	}
	server := grpc.NewServer()
	c.RegisterFunc(server)

	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println(config.C.Server.Name + " 启动GPC失败")

	}
	log.Println(config.C.Server.Name + " GRPC启动成功..." + config.C.Server.Addr)
	err = server.Serve(lis)
	if err != nil {
		log.Println(config.C.Server.Name+" server started error", err)
		return
	}
}
