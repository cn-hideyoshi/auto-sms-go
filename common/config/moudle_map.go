package config

const (
	ModuleDb = iota
	ModuleEtcd
	ModuleGrpc
	ModuleRedis
	ModuleHttp
)

const (
	ReadAll = iota + 1000
	ReadHttp
	ReadGrpc
)

var httpModule [4]int = [4]int{ModuleDb, ModuleEtcd, ModuleHttp, ModuleRedis}
var grpcModule [4]int = [4]int{ModuleDb, ModuleEtcd, ModuleGrpc, ModuleRedis}
var allModule [5]int = [5]int{ModuleDb, ModuleEtcd, ModuleGrpc, ModuleRedis, ModuleHttp}
