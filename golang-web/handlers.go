package main

import (
	"net/http"
)

// viewHandler handle http.Request to "/view/" route and
// render the wiki text/html file to the client.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	var (
		title = r.URL.Path[len("/view/"):]
		page, err = loadPage(title)
	)

	if title == "" {
		resBadRequest(w)
		return
	} else if err != nil {
		http.NotFound(w, r)
		return
	}

	renderTemplate(w, "view", page)
}

// editHandler handle http.Request to "/edit/" route and
// used to edit the wiki from the client page site.
func editHandler(w http.ResponseWriter, r *http.Request) {
	var (
		title = r.URL.Path[len("/edit/"):]
		page, err = loadPage(title)
	)

	if title == "" {
		resBadRequest(w)
		return
	} else if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, "edit", page)
}
