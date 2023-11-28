package config

import (
	"blog.hideyoshi.top/common/config"
)

var C = InitConfig()

type GatewayConfig struct {
	config.Config
}

func InitConfig() *GatewayConfig {
	ViperConfig := config.NewConfig()
	c := &GatewayConfig{
		Config: ViperConfig,
	}
	c.ReadConfig([]int{config.ModuleHttp, config.ModuleEtcd, config.ModuleGrpc})
	return c
}
