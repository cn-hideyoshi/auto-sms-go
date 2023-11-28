package config

type GrpcConfig struct {
	Name string
	Addr string
}

func (c *Config) readGrpcConfig() {
	c.Grpc = &GrpcConfig{
		Name: c.Viper.GetString("grpc.name"),
		Addr: c.Viper.GetString("grpc.addr"),
	}
}
