package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mashiike/surls/usecase"
)

//ShortcutController is default controller
type ShortcutController struct {
	u *usecase.Usecase
}

//NewShortcutController constructs ShortcutController with Usecase
func newShortcutController(u *usecase.Usecase) *ShortcutController {
	c := &ShortcutController{
		u: u,
	}
	return c
}

func (c *ShortcutController) RegisterRoute(r *mux.Router) {
	r.HandleFunc("/{shortcut_id:.+}", c.GetShortcut).Methods("GET")
}

func (c *ShortcutController) GetShortcut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["shortcut_id"]
	if !ok {
		http.NotFound(w, r)
		return
	}
	input := &usecase.GetShortcutInput{
		ShortcutID: id,
	}
	output, err := c.u.GetShortcut(r.Context(), input)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d %s\n", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	http.Redirect(w, r, output.RedirectURL.String(), 301)
}
