package config

import (
	"blog.hideyoshi.top/common/config"
)

var C = InitConfig()

type UserConfig struct {
	config.Config
	Grpc *GrpcConfig
	Etcd *EtcdConfig
}

func InitConfig() *UserConfig {
	ViperConfig := config.NewConfig()
	userConfig := &UserConfig{
		Config: ViperConfig,
	}
	userConfig.ReadServerConfig()
	userConfig.ReadEtcdConfig()
	return userConfig
}

type GrpcConfig struct {
	Name string
	Addr string
}

func (c *UserConfig) ReadServerConfig() {
	c.Grpc = &GrpcConfig{
		Name: c.Viper.GetString("server.name"),
		Addr: c.Viper.GetString("server.addr"),
	}
}

type EtcdConfig struct {
	Addr string
}

func (c *UserConfig) ReadEtcdConfig() {
	// 服务注册
	c.Etcd = &EtcdConfig{
		Addr: c.Viper.GetString("etcd.addr"),
	}
}
