package rpc

import (
	"blog.hideyoshi.top/common/pkg/discovery"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"blog.hideyoshi.top/gateway/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

type GatewayServer struct {
	GatewayResolver *discovery.EtcdResolver
	CancelFunc      context.CancelFunc
	ctx             context.Context
	ResolverClose   func()

	UserLoginClient    userV1.UserLoginServiceClient
	CompanyLoginClient companyV1.CompanyLoginServiceClient
	CompanyInfoClient  companyV1.CompanyInfoServiceClient
}

func NewGateWayServer() *GatewayServer {
	etcdResolver := discovery.NewResolver([]string{config.C.Etcd.Addr})
	resolver.Register(etcdResolver)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	return &GatewayServer{
		ctx:             ctx,
		CancelFunc:      cancelFunc,
		ResolverClose:   etcdResolver.Close,
		GatewayResolver: etcdResolver,
	}
}

var Server *GatewayServer

func init() {
	Server = NewGateWayServer()
	defer Server.ResolverClose()
	Server.NewRpcClient("user", &Server.UserLoginClient)
	Server.NewRpcClient("company", &Server.CompanyLoginClient)
	Server.NewRpcClient("company", &Server.CompanyInfoClient)
	log.Println("init client success...")
}

func (gs GatewayServer) NewRpcClient(serverName string, client interface{}) {
	conn, err := gs.connectServer(serverName)
	if err != nil {
		log.Fatalln("create grpc dial fail", err)
	}

	switch c := client.(type) {
	case *userV1.UserLoginServiceClient:
		*c = userV1.NewUserLoginServiceClient(conn)
	case *companyV1.CompanyLoginServiceClient:
		*c = companyV1.NewCompanyLoginServiceClient(conn)
	case *companyV1.CompanyInfoServiceClient:
		*c = companyV1.NewCompanyInfoServiceClient(conn)
	default:
		log.Fatalln("un support grpc module")
	}
}

func (gs GatewayServer) connectServer(serverName string) (*grpc.ClientConn, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	target := fmt.Sprintf("%s:/%s", "etcd", serverName)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.DialContext(ctx, target, options...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
