package config

const (
	ModuleDb = iota
	ModuleEtcd
	ModuleGrpc
	ModuleRedis

	ReadAll = 999
)

var moduleEnum [4]int = [4]int{ModuleDb, ModuleEtcd, ModuleGrpc, ModuleRedis}
