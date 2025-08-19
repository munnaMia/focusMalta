package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// file server to serve static files
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./cmd/ui/static")))
	mux.Handle("GET /static/", fileServer)

	mux.Handle("GET /", http.HandlerFunc(app.home))
	mux.Handle("GET /about", http.HandlerFunc(app.about))

	return mux
}
