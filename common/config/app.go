package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Viper *viper.Viper
	Grpc  *GrpcConfig
	Etcd  *EtcdConfig
	Db    *DbConfig
	Redis *RedisConfig
}

func NewConfig() Config {
	config := Config{
		Viper: viper.New(),
	}
	workDir, _ := os.Getwd()
	config.Viper.SetConfigName("app")
	config.Viper.SetConfigType("yaml")
	config.Viper.AddConfigPath(workDir + "/config")
	err := config.Viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

func (c *Config) ReadConfig(modules []int) {
	for _, module := range modules {
		switch module {
		case ModuleDb:
			c.ReadDbConfig()
		case ModuleEtcd:
			c.ReadEtcdConfig()
		case ModuleGrpc:
			c.ReadServerConfig()
		case ModuleRedis:
			c.ReadRedisConfig()
		case ReadAll:
			c.ReadConfig(moduleEnum[:])
		}
	}
}
