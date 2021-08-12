package main

import (
	"fmt"
	
	"golang-greet/greeting"
)

func main() {
	// Get a greeting message and print it.
	message := greeting.Hello("Gladys")
	fmt.Println(message)
}
