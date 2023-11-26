package config

type EtcdConfig struct {
	Addr string
}

func (c *Config) ReadEtcdConfig() {
	// 服务注册
	c.Etcd = &EtcdConfig{
		Addr: c.Viper.GetString("etcd.addr"),
	}
}
