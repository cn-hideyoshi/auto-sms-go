package config

type DbConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Charset  string
	DbName   string
}

func (c *Config) readDbConfig() {
	c.Db = &DbConfig{
		Driver:   c.Viper.GetString("db.driver"),
		Host:     c.Viper.GetString("db.host"),
		Port:     c.Viper.GetInt("db.port"),
		Username: c.Viper.GetString("db.username"),
		Password: c.Viper.GetString("db.password"),
		DbName:   c.Viper.GetString("db.db_name"),
	}
}
