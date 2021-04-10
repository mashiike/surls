// +build wireinject

//go:generate wire gen -output_file_prefix dependency_
package surls

import (
	"net/http"

	"github.com/google/wire"
	"github.com/mashiike/surls/controller"
	"github.com/mashiike/surls/usecase"
	"github.com/mashiike/surls/usecase/stub"
)

func NewServeMux(config *Config) http.Handler {
	switch {
	case config.UseStub:
		return newStubServMux(config)
	}
	return nil
}

var commonSet = wire.NewSet(
	controller.NewServeMux,
	getControllerConfig,
	newUsecase,
)
var stubSet = wire.NewSet(
	commonSet,
	stub.NewGetShortcutInteractor,
)

func getControllerConfig(conf *Config) *controller.Config {
	return conf.Route
}

func newStubServMux(conf *Config) http.Handler {
	wire.Build(stubSet)
	return nil
}

func newUsecase(
	getShortcut usecase.GetShortcutInteractor,
) *usecase.Usecase {
	return &usecase.Usecase{
		GetShortcutInteractor: getShortcut,
	}
}
