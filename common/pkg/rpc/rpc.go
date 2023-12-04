package rpc

import (
	"blog.hideyoshi.top/common/config"
	"blog.hideyoshi.top/common/pkg/discovery"
	companyV1 "blog.hideyoshi.top/common/pkg/service/company.v1"
	userV1 "blog.hideyoshi.top/common/pkg/service/user.v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"sync"
	"time"
)

type CommonDiscoveryServer struct {
	GatewayResolver *discovery.EtcdResolver
	CancelFunc      context.CancelFunc
	ctx             context.Context
	ResolverClose   func()

	connMap map[string]*grpc.ClientConn
	mapLock sync.Mutex

	UserLoginClient      userV1.UserLoginServiceClient
	CompanyLoginClient   companyV1.CompanyLoginServiceClient
	CompanyInfoClient    companyV1.CompanyInfoServiceClient
	DepartmentInfoClient companyV1.DepartmentInfoServiceClient
}

func NewCommonDiscoveryServer(config *config.EtcdConfig) *CommonDiscoveryServer {
	etcdResolver := discovery.NewResolver([]string{config.Addr})
	resolver.Register(etcdResolver)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	return &CommonDiscoveryServer{
		ctx:             ctx,
		CancelFunc:      cancelFunc,
		ResolverClose:   etcdResolver.Close,
		connMap:         make(map[string]*grpc.ClientConn),
		GatewayResolver: etcdResolver,
	}
}

func (gs *CommonDiscoveryServer) NewRpcClient(serverName string, client interface{}) {
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
	case *companyV1.DepartmentInfoServiceClient:
		*c = companyV1.NewDepartmentInfoServiceClient(conn)
	default:
		log.Fatalln("not support the grpc module")
	}
}

func (gs *CommonDiscoveryServer) connectServer(serverName string) (*grpc.ClientConn, error) {
	conn, ok := gs.connMap[serverName]
	if ok {
		return conn, nil
	}
	target := fmt.Sprintf("%s:///%s", "etcd", serverName)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	conn, err := grpc.DialContext(timeout, target, options...)
	if err != nil {
		return nil, err
	}
	gs.mapLock.Lock()
	gs.connMap[serverName] = conn
	gs.mapLock.Unlock()
	return conn, nil
}
