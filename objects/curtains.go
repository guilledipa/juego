package objects

import (
	_ "image/png"
	"math"

	"github.com/guilledipa/juego/assets"
	"github.com/guilledipa/juego/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type curtains struct {
	top     string
	lateral string
}

func (c *curtains) Draw(target *ebiten.Image) error {
	targetWidth, _ := target.Size()

	// Lateral
	latImg, err := utils.GetImage(c.lateral, assets.Stall)
	if err != nil {
		return err
	}
	op := &ebiten.DrawImageOptions{}
	target.DrawImage(latImg, op)
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(targetWidth), float64(0))
	target.DrawImage(latImg, op)

	// Top
	topImg, err := utils.GetImage(c.top, assets.Stall)
	if err != nil {
		return err
	}
	topImgW, _ := topImg.Size()
	xTopInstances := int(math.Ceil(float64(targetWidth) / float64(topImgW)))
	for i := 0; i < xTopInstances; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i*topImgW), float64(0))
		target.DrawImage(topImg, op)
	}
	return nil
}

func NewCurtains(lateral, top string) Object {
	return &curtains{
		top:     top,
		lateral: lateral,
	}
}
