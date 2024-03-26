package stbz

// Config Config
type Config struct {
	AccessKey string
	SecretKey string
	Host      string
}

// NewConfig NewConfig
func NewConfig(accessKey string, secretKey string, host string) *Config {
	config := new(Config)
	config.AccessKey = accessKey
	config.SecretKey = secretKey
	config.Host = host
	return config
}
