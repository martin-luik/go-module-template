package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(composition *Composition) http.Handler {
	rootRouter := mux.NewRouter()
	apiRouter := rootRouter.PathPrefix("/api").Subrouter()

	composition.ExampleController.AppendToRouter(apiRouter)

	return rootRouter
}
