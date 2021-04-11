package surls

import (
	"github.com/mashiike/surls/controller"
	"github.com/mashiike/surls/entity"
)

//Config is surls application configure.
type Config struct {
	UseStub bool
	Route   *controller.Config
	Core    *entity.Config
}

func NewDefaultConfig() *Config {
	return &Config{
		UseStub: false,
		Route:   controller.NewDefaultConfig(),
		Core:    entity.NewDefaultConfig(),
	}
}
