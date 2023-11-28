package config

const (
	ModuleDb = iota
	ModuleEtcd
	ModuleGrpc
	ModuleRedis
	ModuleHttp
)

var HttpModules [4]int = [4]int{ModuleDb, ModuleEtcd, ModuleHttp, ModuleRedis}
var GrpcModules [4]int = [4]int{ModuleDb, ModuleEtcd, ModuleGrpc, ModuleRedis}
var AllModules [5]int = [5]int{ModuleDb, ModuleEtcd, ModuleGrpc, ModuleRedis, ModuleHttp}
