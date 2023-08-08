package main

import (
	"fmt"
	"log"

	"example.com/greeting"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	message, err := greeting.Hello("G")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"A", "H"}
	messages, err := greeting.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
