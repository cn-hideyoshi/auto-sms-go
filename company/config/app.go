package config

import (
	"blog.hideyoshi.top/common/config"
)

var C = InitConfig()

type CompanyConfig struct {
	config.Config
	//Grpc  *GrpcConfig
	//Etcd  *EtcdConfig
	//Db    *DbConfig
	//Redis *RedisConfig
}

func InitConfig() *CompanyConfig {
	ViperConfig := config.NewConfig()
	companyConfig := &CompanyConfig{
		Config: ViperConfig,
	}
	modules := []int{config.ReadAll}
	companyConfig.ReadConfig(modules)
	return companyConfig
}
