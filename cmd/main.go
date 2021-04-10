package main

import (
	"log"
	"net/http"
	"training/internal/app"
)

func main() {
	composition := app.NewComposition()
	router := app.NewRouter(composition)

	log.Fatal(http.ListenAndServe(":8080", router))
}
