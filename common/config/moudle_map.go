package config

const (
	ModuleDb = iota
	ModuleEtcd
	ModuleGrpc
	ModuleRedis
	ModuleHttp
	ModuleAmqp
	ModuleSms
)

var GrpcModules = [4]int{ModuleDb, ModuleEtcd, ModuleGrpc, ModuleRedis}
