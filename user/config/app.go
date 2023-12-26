package config

import (
	"blog.hideyoshi.top/common/config"
)

type UserConfig struct {
	config.Config
}

var C = InitConfig()

func InitConfig() *UserConfig {
	ViperConfig := config.NewConfig("user")
	userConfig := &UserConfig{
		Config: ViperConfig,
	}

	userConfig.ReadConfig(config.GrpcModules[:])
	return userConfig
}
