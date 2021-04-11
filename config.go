package surls

import (
	"github.com/mashiike/surls/controller"
	"github.com/mashiike/surls/entity"
)

//go:generate stringer -type=GatewayType

type GatewayType int

const (
	Inmem GatewayType = iota + 1
)

//Config is surls application configure.
type Config struct {
	UseStub bool
	Route   *controller.Config
	Core    *entity.Config
	Gateway GatewayType
}

func NewDefaultConfig() *Config {
	return &Config{
		UseStub: false,
		Route:   controller.NewDefaultConfig(),
		Core:    entity.NewDefaultConfig(),
		Gateway: Inmem,
	}
}
