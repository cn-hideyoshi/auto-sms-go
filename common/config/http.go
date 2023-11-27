package config

type HttpConfig struct {
	Name string
	Addr string
}

func (c *Config) ReadHttpConfig() {
	c.Http = &HttpConfig{
		Name: c.Viper.GetString("server.name"),
		Addr: c.Viper.GetString("server.addr"),
	}
}
