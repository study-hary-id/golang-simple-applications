package main

import (
	"fmt"
	"log"

	"golang-greet/greeting"
)

func main() {
	// Set properties of the predefined logger, including the log entry prefix
	// and a flag to disable printing the time, source file and line number.
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	names := []string{"Darrin", "Gladys", "Samantha"}

	// Request a greeting message and check in if error exists
	message, error := greeting.Hellos(names)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
