package controller

type Config struct {
	DefaultRoute string
	APIRoute     string
}

func NewDefaultConfig() *Config {
	return &Config{
		DefaultRoute: "/",
		APIRoute:     "/api/v1",
	}
}
