package tiropato

import (
	"log"

	"github.com/guilledipa/juego/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

type Game struct {
	objects []objects.Object
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		if err := o.Draw(screen); err != nil {
			log.Fatalf("Draw(): %v", err)
		}
	}
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
	g := &Game{
		objects: []objects.Object{
			objects.NewBackground("bg_green.png"),
			objects.NewDesk("bg_wood.png"),
			objects.NewCurtains("curtains.png", "curtain_straight.png"),
		},
	}
	return g
}
