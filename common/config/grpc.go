package config

type GrpcConfig struct {
	Name string
	Addr string
}

func (c *Config) ReadServerConfig() {
	c.Grpc = &GrpcConfig{
		Name: c.Viper.GetString("server.name"),
		Addr: c.Viper.GetString("server.addr"),
	}
}
