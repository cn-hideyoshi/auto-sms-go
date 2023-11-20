package config

import (
	"blog.hideyoshi.top/common/config"
)

var C = InitConfig()

type CompanyConfig struct {
	config.Config
	Grpc  *GrpcConfig
	Etcd  *EtcdConfig
	Db    *DbConfig
	Redis *RedisConfig
}

func InitConfig() *CompanyConfig {
	ViperConfig := config.NewConfig()
	userConfig := &CompanyConfig{
		Config: ViperConfig,
	}
	userConfig.ReadServerConfig()
	userConfig.ReadEtcdConfig()
	userConfig.ReadDbConfig()
	userConfig.ReadRedisConfig()
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

type RedisConfig struct {
	Host string
	Port uint
	Pass string
	DB   uint
}

func (c *CompanyConfig) ReadRedisConfig() {
	c.Redis = &RedisConfig{
		c.Viper.GetString("redis.host"),
		c.Viper.GetUint("redis.port"),
		c.Viper.GetString("redis.password"),
		c.Viper.GetUint("redis.db"),
	}
}
