package gamelibs


import (
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)



type Game struct {
	touches []ebiten.TouchID
	count   int
}



func getGameTitle() string {
	return "Grow Player - Where the wild thornberries are! (2108)"
}


var brushImage *ebiten.Image

var canvasImage = ebiten.NewImage()

func (g *Game) Update() error {
	drawn := false
	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.paint(canvasImage)
	}
}

func (g *Game) paint(canvas *ebiten.Image, x, y int) {

}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return settings.Game.Width, settings.Game.Height
}

func Run() Game {
	settings := Setup()
	ebiten.SetWindowSize(settings.Game.Width, settings.Game.Height)
	ebiten.SetWindowTitle(getGameTitle())

	var gameInstance Game
	err := ebiten.RunGame(&gameInstance)
	if err != nil {
		processError(err)
	}

	return gameInstance
}