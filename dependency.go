// +build wireinject

//go:generate wire gen -output_file_prefix dependency_
package surls

import (
	"net/http"

	"github.com/google/wire"
	"github.com/mashiike/surls/controller"
	"github.com/mashiike/surls/entity"
	"github.com/mashiike/surls/gateway/inmem"
	"github.com/mashiike/surls/usecase"
	"github.com/mashiike/surls/usecase/stub"
)

var commonSet = wire.NewSet(
	controller.NewServeMux,
	getControllerConfig,
	getEntityConfig,
	entity.NewFactory,
	newUsecase,
	newShortcutRepository,
)
var stubSet = wire.NewSet(
	commonSet,
	stub.NewGetShortcutInteractor,
	stub.NewCreateShortcutInteractor,
)

var prodSet = wire.NewSet(
	commonSet,
	usecase.NewGetShortcutInteractor,
	usecase.NewCreateShortcutInteractor,
)

func NewServeMux(config *Config) http.Handler {
	if config.UseStub {
		return newStubServeMux(config)
	}
	return newProdServeMux(config)
}

func getControllerConfig(conf *Config) *controller.Config {
	return conf.Route
}

func getEntityConfig(conf *Config) *entity.Config {
	return conf.Core
}

func newStubServeMux(conf *Config) http.Handler {
	wire.Build(stubSet)
	return nil
}

func newProdServeMux(conf *Config) http.Handler {
	wire.Build(prodSet)
	return nil
}

func newUsecase(
	getShortcut usecase.GetShortcutBoundary,
	createShortcut usecase.CreateShortcutBoundary,
) *usecase.Usecase {
	return &usecase.Usecase{
		GetShortcutBoundary:    getShortcut,
		CreateShortcutBoundary: createShortcut,
	}
}

func newShortcutRepository(conf *Config) usecase.ShortcutRepository {
	switch conf.Gateway {
	case Inmem:
		return inmem.NewShortcutRepository()
	}
	return nil
}
