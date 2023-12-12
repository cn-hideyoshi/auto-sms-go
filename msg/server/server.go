package server

import (
	msgV1 "blog.hideyoshi.top/common/pkg/service/msg.v1"
	"blog.hideyoshi.top/common/server"
	"blog.hideyoshi.top/msg/config"
	"blog.hideyoshi.top/msg/internal/amqp"
	"blog.hideyoshi.top/msg/internal/crontab"
	"blog.hideyoshi.top/msg/internal/service"
	"google.golang.org/grpc"
)

func Start() {
	cb := crontab.NewCrontab()
	cb.StartD()
	queue := amqp.SmsQueue{}
	queue.Consumer()
	grpcServer := server.NewGrpcServer(config.C.Etcd, config.C.Grpc, cb.StopCrontab)
	grpcServer.RegisterEtcd()
	c := server.GrpcConfig{
		Addr: config.C.Grpc.Addr,
		RegisterFunc: func(server *grpc.Server) {
			msgV1.RegisterMsgGroupServiceServer(server, service.NewMsgGroupService())
		},
	}
	grpcServer.RunGrpc(c)
}
