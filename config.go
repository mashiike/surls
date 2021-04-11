package surls

import "github.com/mashiike/surls/controller"

//Config is surls application configure.
type Config struct {
	UseStub bool
	Route   *controller.Config
}

func NewDefaultConfig() *Config {
	return &Config{
		UseStub: false,
		Route:   controller.NewDefaultConfig(),
	}
}
