package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mashiike/surls/usecase"
)

func NewServeMux(conf *Config, u *usecase.Usecase) http.Handler {
	r := mux.NewRouter()
	apic := newAPIController(u)
	apic.RegisterRoute(r.PathPrefix(conf.APIRoute).Subrouter())
	sc := newShortcutController(u)
	sc.RegisterRoute(r.PathPrefix(conf.DefaultRoute).Subrouter())
	return r
}
