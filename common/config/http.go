package config

type HttpConfig struct {
	Name string
	Addr string
}

func (c *Config) readHttpConfig() {
	c.Http = &HttpConfig{
		Name: c.Viper.GetString("server.name"),
		Addr: c.Viper.GetString("server.addr"),
	}
}
