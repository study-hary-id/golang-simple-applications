package main

import (
	"html/template"
	"net/http"
)

// resBadRequest serve response with Bad Request format.
func resBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("400 status bad request"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// renderTemplate used to render template page in html form.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	err := t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
