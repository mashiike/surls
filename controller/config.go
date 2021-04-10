package controller

type Config struct {
	DefaultRoute string
}

func NewDefaultConfig() *Config {
	return &Config{
		DefaultRoute: "/",
	}
}
