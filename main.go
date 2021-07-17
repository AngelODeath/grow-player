package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	helpers "github.com/angelodeath/fynelibs"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	log.SetPrefix("grow-player: ")
	log.SetFlags(0)

	aRandomNumber := helpers.GenerateRandomNumber(100)

	aRandomString, err := helpers.GenerateRandomString(50)
	if err != nil {
		log.Fatal(err)
	}

	luckyNumberMessage := fmt.Sprintf("Your random lucky number:\n  %d\n", aRandomNumber)
	luckyStringMessage := fmt.Sprintf("Your 50 random characters:\n  %v\n", aRandomString)
	fmt.Println(luckyNumberMessage)
	fmt.Println(luckyStringMessage)
}