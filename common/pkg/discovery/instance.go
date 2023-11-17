package discovery

import "google.golang.org/grpc/resolver"

type GrpcServer struct {
	Name string
	Addr string
}

func ResolverExits(resolvers []resolver.Address, resolver resolver.Address) bool {
	for _, data := range resolvers {
		if data.Addr == resolver.Addr {
			return true
		}
	}
	return false
}
