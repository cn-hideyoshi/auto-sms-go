package config

import (
	"blog.hideyoshi.top/common/config"
	"github.com/redis/go-redis/v9"
)

var C = InitConfig()

type CompanyConfig struct {
	config.Config
	Grpc *GrpcConfig
	Etcd *EtcdConfig
	Db   *DbConfig
}

func InitConfig() *CompanyConfig {
	ViperConfig := config.NewConfig()
	userConfig := &CompanyConfig{
		Config: ViperConfig,
	}
	userConfig.ReadServerConfig()
	userConfig.ReadEtcdConfig()
	userConfig.ReadDbConfig()
	return userConfig
}

type GrpcConfig struct {
	Name string
	Addr string
}

func (c *CompanyConfig) ReadServerConfig() {
	c.Grpc = &GrpcConfig{
		Name: c.Viper.GetString("server.name"),
		Addr: c.Viper.GetString("server.addr"),
	}
}

type DbConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Charset  string
	DbName   string
}

func (c *CompanyConfig) ReadDbConfig() {
	c.Db = &DbConfig{
		Driver:   c.Viper.GetString("db.driver"),
		Host:     c.Viper.GetString("db.host"),
		Port:     c.Viper.GetInt("db.port"),
		Username: c.Viper.GetString("db.username"),
		Password: c.Viper.GetString("db.password"),
		DbName:   c.Viper.GetString("db.db_name"),
	}
}

type EtcdConfig struct {
	Addr string
}

func (c *CompanyConfig) ReadEtcdConfig() {
	// 服务注册
	c.Etcd = &EtcdConfig{
		Addr: c.Viper.GetString("etcd.addr"),
	}
}

func (c *CompanyConfig) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     c.Viper.GetString("redis.host") + ":" + c.Viper.GetString("redis.port"),
		Password: c.Viper.GetString("redis.password"),
		DB:       c.Viper.GetInt("redis.db"),
	}
}
