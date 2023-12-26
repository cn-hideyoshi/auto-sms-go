package config

import (
	"blog.hideyoshi.top/common/config"
)

var C = InitConfig()

type CompanyConfig struct {
	config.Config
}

func InitConfig() *CompanyConfig {
	ViperConfig := config.NewConfig("company")
	c := &CompanyConfig{
		Config: ViperConfig,
	}
	c.ReadConfig(config.GrpcModules[:])
	return c
}
