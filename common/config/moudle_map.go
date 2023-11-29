package config

const (
	ModuleDb = iota
	ModuleEtcd
	ModuleGrpc
	ModuleRedis
	ModuleHttp
)

var GrpcModules [4]int = [4]int{ModuleDb, ModuleEtcd, ModuleGrpc, ModuleRedis}
