package config

import (
	"blog.hideyoshi.top/common/config"
)

var C = InitConfig()

type MsgConfig struct {
	config.Config
}

func InitConfig() *MsgConfig {
	ViperConfig := config.NewConfig()
	c := &MsgConfig{
		Config: ViperConfig,
	}
	c.ReadConfig(config.GrpcModules[:])
	return c
}
