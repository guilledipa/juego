package objects

import (
	_ "image/png"
	"math"

	"github.com/guilledipa/juego/assets"
	"github.com/guilledipa/juego/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type desk struct {
	name string
}

func (d *desk) Draw(target *ebiten.Image) error {
	desk, err := utils.GetImage(d.name, assets.Stall)
	if err != nil {
		return err
	}

	targetWidth, targetHeight := target.Size()
	deskWidth, deskHeight := desk.Size()
	xInstances := int(math.Ceil(float64(targetWidth) / float64(deskWidth)))
	yPosition := targetHeight - deskHeight

	for j := 0; j < xInstances; j++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(j*deskWidth), float64(yPosition))
		target.DrawImage(desk, op)
	}
	return nil
}

func NewDesk(img string) Object {
	return &desk{
		name: img,
	}
}
