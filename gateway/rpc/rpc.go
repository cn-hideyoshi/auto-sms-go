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
	"sync"
	"time"
)

// GatewayServer represents the gRPC gateway server that manages connections to various gRPC services.
type GatewayServer struct {
	GatewayResolver *discovery.EtcdResolver
	CancelFunc      context.CancelFunc
	ctx             context.Context
	ResolverClose   func()

	connMap map[string]*grpc.ClientConn
	mapLock sync.Mutex

	UserLoginClient    userV1.UserLoginServiceClient
	CompanyLoginClient companyV1.CompanyLoginServiceClient
	CompanyInfoClient  companyV1.CompanyInfoServiceClient
}

// NewGateWayServer creates a new instance of the GatewayServer.
func NewGateWayServer() *GatewayServer {
	etcdResolver := discovery.NewResolver([]string{config.C.Etcd.Addr})
	resolver.Register(etcdResolver)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	return &GatewayServer{
		ctx:             ctx,
		CancelFunc:      cancelFunc,
		ResolverClose:   etcdResolver.Close,
		connMap:         make(map[string]*grpc.ClientConn),
		GatewayResolver: etcdResolver,
	}
}

// Server is a singleton instance of GatewayServer.
var Server *GatewayServer

// init initializes the GatewayServer and gRPC clients.
func init() {
	Server = NewGateWayServer()

	Server.NewRpcClient("user", &Server.UserLoginClient)
	Server.NewRpcClient("company", &Server.CompanyLoginClient)
	Server.NewRpcClient("company", &Server.CompanyInfoClient)
	defer Server.ResolverClose()
	log.Println("init client success...")
}

// NewRpcClient creates a new gRPC client for the specified server name and assigns it to the provided client interface.
func (gs *GatewayServer) NewRpcClient(serverName string, client interface{}) {
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
		log.Fatalln("not support the grpc module")
	}
}

// connectServer establishes a connection to the specified gRPC server and returns the client connection.
func (gs *GatewayServer) connectServer(serverName string) (*grpc.ClientConn, error) {
	conn, ok := gs.connMap[serverName]
	if ok {
		return conn, nil
	} else {
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		target := fmt.Sprintf("%s:///%s", "etcd", serverName)
		options := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}
		conn, err := grpc.DialContext(ctx, target, options...)
		if err != nil {
			return nil, err
		}
		gs.mapLock.Lock()
		gs.connMap[serverName] = conn
		gs.mapLock.Unlock()
		return conn, nil
	}
}
