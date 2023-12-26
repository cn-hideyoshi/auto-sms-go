package server

import (
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/common/server"
	"blog.hideyoshi.top/user/config"
	"blog.hideyoshi.top/user/internal/service"
	"google.golang.org/grpc"

	_ "blog.hideyoshi.top/user/rpc"
)

func Start() {
	grpcServer := server.NewGrpcServer(config.C.Etcd, config.C.Grpc)
	grpcServer.RegisterEtcd()
	c := server.GrpcConfig{
		Addr: config.C.Grpc.Addr,
		RegisterFunc: func(server *grpc.Server) {
			userV1.RegisterUserPhoneServiceServer(server, service.NewUserPhoneService())
			userV1.RegisterUserInfoServiceServer(server, service.NewUserInfoService())
		},
	}
	grpcServer.RunGrpc(c)
}
