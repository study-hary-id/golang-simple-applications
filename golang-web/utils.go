package main

import (
	"errors"
	"html/template"
	"net/http"
	"regexp"
)

// templates initialize available html templates.
var templates = template.Must(template.ParseFiles(
"view.html", "edit.html",
))

// validPath initializes regexp rule to get a valid URL.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// resBadRequest serves response with Status Bad Request format.
func resBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("400 status bad request"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//resInternalServerError serves response with Status Internal Server Error format.
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

// getTitle returns the requested wiki title and error if invalid.
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	path := validPath.FindStringSubmatch(r.URL.Path)
	if path == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}
	return path[2], nil
}
