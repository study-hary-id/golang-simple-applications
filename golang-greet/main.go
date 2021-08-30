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
	message, err := greeting.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
