package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mashiike/surls/usecase"
)

//APIController is REST API controller
type APIController struct {
	u *usecase.Usecase
}

//newAPIController constructs APIController with Usecase
func NewAPIController(u *usecase.Usecase) *APIController {
	c := &APIController{
		u: u,
	}
	return c
}

func (c *APIController) RegisterRoute(r *mux.Router) {
	r.HandleFunc("/urls", c.CreateShortcut).Methods("POST")
	r.HandleFunc("/*", c.CatchAll)
}

func (c *APIController) CreateShortcut(w http.ResponseWriter, r *http.Request) {
	jw := newJSONResponseWriter(w)
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var input usecase.CreateShortcutInput
	if err := dec.Decode(&input); err != nil {
		log.Print(err)
		jw.WriteError(http.StatusBadRequest, err)
		return
	}
	output, err := c.u.CreateShortcut(r.Context(), &input)
	if err != nil {
		log.Print(err)
		jw.WriteError(http.StatusInternalServerError, err)
		return
	}
	jw.WriteHeader(http.StatusOK)
	jw.Header().Add("Content-Type", "application/json")
	jw.WriteHeader(http.StatusOK)
	jw.WriteBody(map[string]interface{}{
		"shortcut_id": output.ShortcutID,
		"long_url":    output.LongURL,
	})
}

func (c *APIController) CatchAll(w http.ResponseWriter, r *http.Request) {
	jw := newJSONResponseWriter(w)
	jw.WriteError(http.StatusNotFound, fmt.Errorf("%s not found", r.URL.Path))
}
