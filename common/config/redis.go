package config

type RedisConfig struct {
	Host   string
	Port   uint
	Prefix string
	Pass   string
	DB     uint
}

func (c *Config) readRedisConfig() {
	c.Redis = &RedisConfig{
		c.Viper.GetString("redis.host"),
		c.Viper.GetUint("redis.port"),
		c.Viper.GetString("redis.prefix"),
		c.Viper.GetString("redis.password"),
		c.Viper.GetUint("redis.db"),
	}
}
