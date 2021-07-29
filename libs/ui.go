package fynelibs

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Launcher struct {
	app fyne.App
	win fyne.Window
}

type MenuButton struct {
	text string
	action func()
}

type Menu struct {
	buttons []MenuButton
}

func BuildMenu(menu Menu) *fyne.Container {
	menuWrap := container.NewVBox()
	for _, menuBtn := range menu.buttons {
		_btn := widget.NewButton(menuBtn.text, menuBtn.action)
		menuWrap.Add(_btn)
	}
	return menuWrap
}

func AnyButtonClicked() {
	//outputLabel.SetText(outputLabelPrefixStr + "Showing about window...")
	//AppShowAnother(mainApp)
	fmt.Println("A button was clicked.")
}

func ShowMainMenu(launcher Launcher) {

	var mainMenu Menu
	mainMenu.buttons = []MenuButton {
		{"New game", AnyButtonClicked},
		{"Exit", AnyButtonClicked},
	}
	menuContainer := BuildMenu(mainMenu)
	launcher.win.SetContent(menuContainer)
}

func ShowLauncher() Launcher {
	_app := app.New()
	_win := _app.NewWindow("Grow player")


	gameLauncher := Launcher{_app, _win}
	ShowMainMenu(gameLauncher)

	_win.Show()
	_app.Run()

	return gameLauncher
}