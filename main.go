package main

import (
	"log"

	"tutorial/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Lshortfile) // 16

	message := greetings.Hello("Gladys")
	log.Println(message)
}
