package main

import (
	"log"

	"tutorial/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Lshortfile) // 16

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(message)
}
