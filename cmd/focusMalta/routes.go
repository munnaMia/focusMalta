package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(app.home))
	mux.Handle("GET /about", http.HandlerFunc(app.about))

	return mux
}
