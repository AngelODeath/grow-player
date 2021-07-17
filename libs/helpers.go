package fynelibs

import (
	"errors"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	//"github.com/pelletier/go-toml/v2"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func AppSetup(appName string) fyne.App {
	fmt.Printf("Launching \"%s\" ...\n", appName)
	return app.New()
}

func AppCleanup() {
	fmt.Println("Exiting...")
}

func AppShowAnother(a fyne.App) {
	time.Sleep(time.Second * 5)
	win := a.NewWindow("About")
	win.SetContent(widget.NewLabel("... 5 seconds later"))
	win.Resize(fyne.NewSize(320, 240))

	win.Show()

	time.Sleep(time.Second * 2)
	win.Close()
}

func AppMainWindow() {

}

func AppRun() {
	const outputLabelPrefixStr string = "Output: "
	const appName string = "Mine Fyne GUI"

	mainApp := AppSetup(appName)

	mainWindow := mainApp.NewWindow(appName)

	headerLabel := widget.NewLabel("* Welcome to my \"\" *")
	outputLabel := widget.NewLabel(outputLabelPrefixStr)
	footerLabel := widget.NewLabel("* copyright 2021 - Gabri Botha / School of Death *")

	exampleButton := widget.NewButton("Test!", func() {
		outputLabel.SetText(outputLabelPrefixStr + "Showing about window...")
		AppShowAnother(mainApp)
	})
	rootContainer := container.NewVBox(headerLabel, outputLabel, exampleButton, footerLabel)
	mainWindow.SetContent(rootContainer)

	mainWindow.Show()
	mainApp.Run()

	AppCleanup()
}
