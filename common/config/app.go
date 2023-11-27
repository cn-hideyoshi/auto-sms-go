package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Viper *viper.Viper
	//Modules []Configurable
	Grpc  *GrpcConfig
	Etcd  *EtcdConfig
	Db    *DbConfig
	Redis *RedisConfig
	Http  *HttpConfig
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
			c.ReadGrpcConfig()
		case ModuleHttp:
			c.ReadHttpConfig()
		case ModuleRedis:
			c.ReadRedisConfig()
		case ReadHttp:
			c.ReadConfig(httpModule[:])
		case ReadGrpc:
			c.ReadConfig(grpcModule[:])
		case ReadAll:
			c.ReadConfig(allModule[:])
		}
	}
}
