package tiropato

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Run() error {
	return ebiten.RunGame(g)
}

func NewGame() *Game {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Tiro al Pato!")
	g := &Game{}
	return g
}
