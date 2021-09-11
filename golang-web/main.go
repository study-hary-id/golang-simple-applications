package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save/", saveHandler)

	fmt.Println("Server listening at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
