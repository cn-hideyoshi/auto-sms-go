package config

type SmsConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	Url             string
}

func (c *Config) readSmsConfig() {
	c.Sms = &SmsConfig{
		c.Viper.GetString("sms.accessKey"),
		c.Viper.GetString("sms.accessKeySecret"),
		c.Viper.GetString("sms.url"),
	}
}
