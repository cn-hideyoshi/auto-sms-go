package config

type AmqpConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
}

func (c *Config) readAmqpConfig() {
	c.Amqp = &AmqpConfig{
		Driver:   c.Viper.GetString("amqp.driver"),
		Host:     c.Viper.GetString("amqp.host"),
		Port:     c.Viper.GetInt("amqp.port"),
		Username: c.Viper.GetString("amqp.username"),
		Password: c.Viper.GetString("amqp.password"),
	}
}
