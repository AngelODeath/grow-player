// Copyright 2021 Gabri Botha
// 23-07-2021

package main

import (
	"fmt"
	"log"

	helpers "github.com/angelodeath/fynelibs"
)

func init() {
	log.SetPrefix("grow-player: ")
	log.SetFlags(0)
}

func main() {
	settings := helpers.Setup()
	settingsStr := "Settings:\n"
	settingsStr += fmt.Sprintf("  Resolution: %vx%v\n", settings.Game.Width, settings.Game.Height)
	settingsStr += fmt.Sprintf("  Fullscreen: %v\n", settings.Game.Fullscreen)
	settingsStr += fmt.Sprintf("  Save directory: \"%v\"\n", settings.Game.SavePath)
	fmt.Println(settingsStr)

	aRandomNumber := helpers.GenerateRandomNumber(100)
	luckyNumberMessage := fmt.Sprintf("Your random lucky number:\n  %d\n", aRandomNumber)
	fmt.Println(luckyNumberMessage)

	aRandomString := helpers.GenerateRandomString(50)
	luckyStringMessage := fmt.Sprintf("Your 50 random characters:\n  %v\n", aRandomString)
	fmt.Println(luckyStringMessage)

	myIP, err := helpers.LookupMyIp()
	if err != nil {
		log.Fatal(err)
	}
	publicIpMessage := fmt.Sprintf("Your public  IP address:\n  %v\n", myIP)
	fmt.Println(publicIpMessage)

	ui := helpers.ShowLauncher()
	fmt.Println(ui)
}