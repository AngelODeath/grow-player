package main

import (
	"fmt"
	"log"

	helpers "github.com/angelodeath/fynelibs"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	greetMsg, err := helpers.Hello("Jimbo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greetMsg)
	helpers.AppRun()
}
