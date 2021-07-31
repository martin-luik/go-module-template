package example

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

func (this Controller) AppendToRouter(apiRouter *mux.Router) {
	messageRouter := apiRouter.Path("/example").Subrouter()
	messageRouter.Methods(http.MethodGet).HandlerFunc(this.get)
}

func (this Controller) get(w http.ResponseWriter, _ *http.Request) {
	examples := []ExampleJson{
		{Title: "Hello", Description: "World"},
		{Title: "Hello2", Description: "World2"},
	}
	err := json.NewEncoder(w).Encode(examples)
	if err != nil {
		return
	}
}
