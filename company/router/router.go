package router

import (
	"blog.hideyoshi.top/common/pkg/discovery"
	loginServiceV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	"blog.hideyoshi.top/company/config"
	"blog.hideyoshi.top/company/pkg/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterEtcd() *discovery.EtcdRegister {
	etcdAddr := []string{config.C.Etcd.Addr}
	register := discovery.NewRegister(etcdAddr)

	server := discovery.GrpcServer{
		Name: config.C.Grpc.Name,
		Addr: config.C.Grpc.Addr,
	}

	_, err := register.Register(server, 10)
	if err != nil {
		log.Fatalln("register etcd fail :", err)
	}
	return register
}

func RegisterGrpc() {
	c := GrpcConfig{
		Addr: config.C.Grpc.Addr,
		RegisterFunc: func(server *grpc.Server) {
			loginServiceV1.RegisterCompanyLoginServiceServer(server, rpc.NewCompanyLoginService())
			loginServiceV1.RegisterCompanyInfoServiceServer(server, rpc.NewCompanyInfoService())
		},
	}
	server := grpc.NewServer()
	c.RegisterFunc(server)

	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println(config.C.Grpc.Name + " 启动GPC失败")

	}
	log.Println(config.C.Grpc.Name + " GRPC启动成功..." + config.C.Grpc.Addr)
	err = server.Serve(lis)
	if err != nil {
		log.Println(config.C.Grpc.Name+" server started error", err)
		return
	}
}
