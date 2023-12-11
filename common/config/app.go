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
	Http  *HttpConfig
	Etcd  *EtcdConfig
	Db    *DbConfig
	Redis *RedisConfig
	Amqp  *AmqpConfig
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
			c.readDbConfig()
		case ModuleEtcd:
			c.readEtcdConfig()
		case ModuleGrpc:
			c.readGrpcConfig()
		case ModuleHttp:
			c.readHttpConfig()
		case ModuleRedis:
			c.readRedisConfig()
		case ModuleAmqp:
			c.readAmqpConfig()
		}
	}
}
