package main

import (
	"fmt"
	"net/http"
)

// viewHandler handle http.Request to "/view/" route and
// response the text/html file to the client.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	var (
		title = r.URL.Path[len("/view/"):]
		page, err = loadPage(title)
	)

	if title == "" {
		// TODO: Create a function with name resBadRequest
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("400 status bad request"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	} else if err != nil {
		http.NotFound(w, r)
		return
	}

	_, err = fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
