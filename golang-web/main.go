package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page model to construct the content of a web page.
type Page struct {
	Title string
	Body []byte
}

// save writes file .txt to current local directory and
// return error if there is a problem while writing.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage reads file .txt in current local directory and
// return error if the file does not exist.
func loadPage(title string) (*Page, error) {
	var (
		filename = title + ".txt"
		body, err = ioutil.ReadFile(filename)
	)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// viewHandler handle http.Request to "/view/" route and
// response the text/html file to the client.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	var (
		title = r.URL.Path[len("/view/"):]
		page, _ = loadPage(title)
	)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func main() {
	//savingPage := &Page{Title: "TestPage", Body: []byte(
	//	"This is a sample page to saving and loading",
	//)}
	//savingPage.save()

	//loadingPage, _ := loadPage("TestPage")
	//fmt.Println(string(loadingPage.Body))

	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
