package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"os"
)

var C = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
	Grpc  *GrpcConfig
}

type GrpcConfig struct {
	Name string
	Addr string
}

type ServerConfig struct {
	Name string
	Addr string
}

func InitConfig() *Config {
	config := &Config{
		viper: viper.New(),
	}
	workDir, _ := os.Getwd()
	config.viper.SetConfigName("app")
	config.viper.SetConfigType("yaml")
	config.viper.AddConfigPath(workDir + "/config")

	err := config.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	config.ReadServerConfig()
	config.ReadGrpcConfig()

	return config
}

func (c *Config) ReadGrpcConfig() {
	serv := &GrpcConfig{}
	serv.Name = c.viper.GetString("server.name")
	serv.Addr = c.viper.GetString("server.addr")
	c.Grpc = serv
}
func (c *Config) ReadServerConfig() {
	serv := &ServerConfig{}
	serv.Name = c.viper.GetString("server.name")
	serv.Addr = c.viper.GetString("server.addr")
	c.SC = serv
}

func (c *Config) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("redis.db"),
	}
}
