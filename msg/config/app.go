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
	modules := config.GrpcModules[:]
	modules = append(modules, config.ModuleAmqp)
	c.ReadConfig(modules)
	return c
}
