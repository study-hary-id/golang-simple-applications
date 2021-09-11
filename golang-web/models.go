package main

import "io/ioutil"

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
