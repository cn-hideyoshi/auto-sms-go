package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Viper *viper.Viper
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
