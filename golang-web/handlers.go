package main

import (
	"net/http"
)

// viewHandler handle http.Request to "/view/" route and
// render the wiki text/html file to the client.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", page)
}

// editHandler handle http.Request to "/edit/" route and
// used to edit the wiki from the client page site.
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, "edit", page)
}

// saveHandler handle http.Request to "/save/" route and
// redirect to "/view/" after saving the wiki.
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	var (
		body = r.FormValue("body")
		page = &Page{Title: title, Body: []byte(body)}
	)
	err = page.save()
	if err != nil {
		resInternalServerError(w, err)
		return
	}

	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}
