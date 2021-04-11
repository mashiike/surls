// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package surls

import (
	"github.com/google/wire"
	"github.com/mashiike/surls/controller"
	"github.com/mashiike/surls/entity"
	"github.com/mashiike/surls/usecase"
	"github.com/mashiike/surls/usecase/stub"
	"net/http"
)

// Injectors from dependency.go:

func newStubServeMux(conf *Config) http.Handler {
	config := getControllerConfig(conf)
	getShortcutBoundary := stub.NewGetShortcutInteractor()
	createShortcutBoundary := stub.NewCreateShortcutInteractor()
	usecase := newUsecase(getShortcutBoundary, createShortcutBoundary)
	handler := controller.NewServeMux(config, usecase)
	return handler
}

func newProdServeMux(conf *Config) http.Handler {
	config := getControllerConfig(conf)
	getShortcutBoundary := stub.NewGetShortcutInteractor()
	entityConfig := getEntityConfig(conf)
	factory := entity.NewFactory(entityConfig)
	createShortcutBoundary := usecase.NewCreateShortcutInteractor(factory)
	usecaseUsecase := newUsecase(getShortcutBoundary, createShortcutBoundary)
	handler := controller.NewServeMux(config, usecaseUsecase)
	return handler
}

// dependency.go:

var commonSet = wire.NewSet(controller.NewServeMux, getControllerConfig,
	getEntityConfig, entity.NewFactory, newUsecase,
)

var stubSet = wire.NewSet(
	commonSet, stub.NewGetShortcutInteractor, stub.NewCreateShortcutInteractor,
)

var prodSet = wire.NewSet(
	commonSet, stub.NewGetShortcutInteractor, usecase.NewCreateShortcutInteractor,
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

func newUsecase(
	getShortcut usecase.GetShortcutBoundary,
	createShortcut usecase.CreateShortcutBoundary,
) *usecase.Usecase {
	return &usecase.Usecase{
		GetShortcutBoundary:    getShortcut,
		CreateShortcutBoundary: createShortcut,
	}
}
