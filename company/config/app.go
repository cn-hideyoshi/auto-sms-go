package config

import (
	"blog.hideyoshi.top/common/config"
)

var C = InitConfig()

type CompanyConfig struct {
	config.Config
}

func InitConfig() *CompanyConfig {
	ViperConfig := config.NewConfig()
	c := &CompanyConfig{
		Config: ViperConfig,
	}
	c.ReadConfig(config.GrpcModules[:])
	return c
}
