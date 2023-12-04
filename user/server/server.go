package server

import (
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/common/server"
	"blog.hideyoshi.top/user/config"
	"blog.hideyoshi.top/user/internal/service"
	"blog.hideyoshi.top/user/rpc"
	"google.golang.org/grpc"
)
import _ "blog.hideyoshi.top/user/rpc"

func Start() {
	grpcServer := server.NewGrpcServer(config.C.Etcd, config.C.Grpc, rpc.Clients.ResolverClose)
	grpcServer.RegisterEtcd()
	c := server.GrpcConfig{
		Addr: config.C.Grpc.Addr,
		RegisterFunc: func(server *grpc.Server) {
			userV1.RegisterUserLoginServiceServer(server, service.NewUserLoginService())
			userV1.RegisterUserInfoServiceServer(server, service.NewUserInfoService())
		},
	}
	grpcServer.RunGrpc(c)
}
