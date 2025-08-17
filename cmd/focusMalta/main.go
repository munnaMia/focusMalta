package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	app := application{}

	// contain  server configurations.
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
