package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mashiike/surls/usecase"
)

func NewServeMux(conf *Config, u *usecase.Usecase) http.Handler {
	r := mux.NewRouter()
	sc := newShortcutController(u)
	sc.RegisterRoute(r.PathPrefix(conf.DefaultRoute).Subrouter())
	r.HandleFunc("/", catchAll)
	return r
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
