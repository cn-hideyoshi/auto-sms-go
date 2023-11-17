package config

import (
	"blog.hideyoshi.top/common/config"
	"github.com/redis/go-redis/v9"
)

var C = InitConfig()

type UserConfig struct {
	config.Config
	Server *ServerConfig
	Etcd   *EtcdConfig
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

type ServerConfig struct {
	Name string
	Addr string
}

func (c *UserConfig) ReadServerConfig() {
	c.Server = &ServerConfig{
		Name: c.Viper.GetString("server.name"),
		Addr: c.Viper.GetString("server.addr"),
	}
}

func (c *UserConfig) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     c.Viper.GetString("redis.host") + ":" + c.Viper.GetString("redis.port"),
		Password: c.Viper.GetString("redis.password"),
		DB:       c.Viper.GetInt("redis.db"),
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
