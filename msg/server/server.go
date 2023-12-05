package server

import (
	msgV1 "blog.hideyoshi.top/common/pkg/service/msg.v1"
	"blog.hideyoshi.top/common/server"
	"blog.hideyoshi.top/msg/config"
	"blog.hideyoshi.top/msg/internal/service"
	"blog.hideyoshi.top/msg/rpc"
	"google.golang.org/grpc"
)

func Start() {
	grpcServer := server.NewGrpcServer(config.C.Etcd, config.C.Grpc, rpc.Clients.ResolverClose)
	grpcServer.RegisterEtcd()
	c := server.GrpcConfig{
		Addr: config.C.Grpc.Addr,
		RegisterFunc: func(server *grpc.Server) {
			msgV1.RegisterMsgGroupServiceServer(server, service.NewMsgGroupService())
		},
	}
	grpcServer.RunGrpc(c)
}
