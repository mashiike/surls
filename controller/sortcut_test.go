package controller_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mashiike/surls/controller"
	"github.com/mashiike/surls/usecase"
	"github.com/mashiike/surls/usecase/stub"
	"github.com/stretchr/testify/assert"
)

func TestShortcutController(t *testing.T) {
	c := controller.NewShortcutController(&usecase.Usecase{
		GetShortcutBoundary: stub.NewGetShortcutInteractor(),
	})
	router := mux.NewRouter()
	c.RegisterRoute(router)
	assert.HTTPRedirect(t, http.HandlerFunc(router.ServeHTTP), "GET", "/test", nil)
}
