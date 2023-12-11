package server

import (
	"blog.hideyoshi.top/common/config"
	"blog.hideyoshi.top/common/pkg/discovery"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GrpcServer struct {
	etcdConfig *config.EtcdConfig
	grpcConfig *config.GrpcConfig
	stopFunc   []func()
}

func NewGrpcServer(etcd *config.EtcdConfig, grpc *config.GrpcConfig, stopFunc ...func()) *GrpcServer {
	return &GrpcServer{
		etcd,
		grpc,
		stopFunc,
	}
}

func (g GrpcServer) RegisterEtcd() *discovery.EtcdRegister {
	etcdAddr := []string{g.etcdConfig.Addr}
	register := discovery.NewRegister(etcdAddr)

	server := discovery.GrpcServer{
		Name: g.grpcConfig.Name,
		Addr: g.grpcConfig.Addr,
	}

	_, err := register.Register(server, 10)
	if err != nil {
		log.Fatalln("register etcd fail :", err)
	}
	return register
}

type GrpcConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}

func (g GrpcServer) RunGrpc(c GrpcConfig) {
	defer log.Printf("%s关闭成功", g.grpcConfig.Name)
	server := grpc.NewServer()
	go func() {
		c.RegisterFunc(server)
		lis, err := net.Listen("tcp", c.Addr)
		if err != nil {
			log.Println(g.grpcConfig.Name + " 启动GPC失败")

		}
		log.Println(g.grpcConfig.Name + " GRPC启动成功..." + g.grpcConfig.Addr)
		err = server.Serve(lis)
		if err != nil {
			log.Println(g.grpcConfig.Name+" server started error", err)
			return
		}
	}()
	quit := make(chan os.Signal)
	// SIGINT 用户发送INTR字符(Ctrl+C)触发 kill -2
	// SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("正在关闭程序%s", g.grpcConfig.Name)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	server.GracefulStop()
	for _, f := range g.stopFunc {
		f()
	}
	select {
	case <-ctx.Done():
		log.Printf("等待%s关闭...", g.grpcConfig.Name)
	}
}
