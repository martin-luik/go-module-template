package main

import (
	"example/internal/app"
	"log"
	"net/http"
)

func main() {
	composition := app.NewComposition()
	router := app.NewRouter(composition)

	log.Fatal(http.ListenAndServe(":8080", router))
}
