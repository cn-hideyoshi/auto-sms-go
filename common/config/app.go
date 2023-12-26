package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
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
	Sms   *SmsConfig
}

func NewConfig(serverName string) Config {
	config := Config{
		Viper: viper.New(),
	}
	workDir, _ := os.Getwd()
	config.Viper.SetConfigName(serverName)
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
		case ModuleSms:
			c.readSmsConfig()
		}
	}
}
