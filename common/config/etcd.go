package config

type EtcdConfig struct {
	Addr string
}

func (c *Config) readEtcdConfig() {
	// 服务注册
	c.Etcd = &EtcdConfig{
		Addr: c.Viper.GetString("etcd.addr"),
	}
}
