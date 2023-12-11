package config

const (
	ModuleDb = iota
	ModuleEtcd
	ModuleGrpc
	ModuleRedis
	ModuleHttp
	ModuleAmqp
)

var GrpcModules = [4]int{ModuleDb, ModuleEtcd, ModuleGrpc, ModuleRedis}
