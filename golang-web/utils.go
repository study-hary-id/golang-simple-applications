package main

import (
	"html/template"
	"net/http"
)

// templates initialize available html templates.
var templates = template.Must(template.ParseFiles(
"view.html", "edit.html",
))

// resBadRequest serve response with Status Bad Request format.
func resBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("400 status bad request"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//resInternalServerError serve response with Status Internal Server Error format.
func resInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// renderTemplate used to render template page in html form.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		resInternalServerError(w, err)
		return
	}
}
