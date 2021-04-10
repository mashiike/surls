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
func newAPIController(u *usecase.Usecase) *APIController {
	c := &APIController{
		u: u,
	}
	return c
}

func (c *APIController) RegisterRoute(r *mux.Router) {
	//r.Headers("Content-Type", "application/json")
	r.HandleFunc("/urls", c.CreateShortcut).Methods("POST")
}

func (c *APIController) CreateShortcut(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var input usecase.CreateShortcutInput
	if err := dec.Decode(&input); err != nil {
		log.Print(err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
			"detail":  err.Error(),
		})
		return
	}
	output, err := c.u.CreateShortcut(r.Context(), &input)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d %s\n", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"shortcut_id": output.ShortcutID,
		"long_url":    output.LongURL,
	})
}
