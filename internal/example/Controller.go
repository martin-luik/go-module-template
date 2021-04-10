package example

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Example struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

func (this Controller) AppendToRouter(apiRouter *mux.Router) {
	messageRouter := apiRouter.Path("/example").Subrouter()
	messageRouter.Methods(http.MethodGet).HandlerFunc(this.exampleGet)
}

func (this Controller) exampleGet(w http.ResponseWriter, r *http.Request) {
	examples := []Example{
		Example{Title: "Hello", Description: "World"},
		Example{Title: "Hello2", Description: "World2"},
	}
	err := json.NewEncoder(w).Encode(examples)
	if err != nil {
		return
	}
}
