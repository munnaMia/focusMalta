package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./cmd/ui/html/base.html",
		"./cmd/ui/html/pages/home.html",
		"./cmd/ui/html/partials/nav.html",
		"./cmd/ui/html/partials/footer.html",
		"./cmd/ui/html/partials/header.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ErrorLog.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.ErrorLog.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GET - /about - handler
func (app *application) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about us"))
}
