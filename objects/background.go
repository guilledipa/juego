package objects

import (
	_ "image/png"
	"math"

	"github.com/guilledipa/juego/assets"
	"github.com/guilledipa/juego/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type background struct {
	name string
}

func (b *background) Draw(target *ebiten.Image) error {
	bg, err := utils.GetImage(b.name, assets.Stall)
	if err != nil {
		return err
	}
	targetWidth, targetHeight := target.Size()
	bgWidth, bgHeight := bg.Size()
	xInstances := int(math.Ceil(float64(targetWidth) / float64(bgWidth)))
	yInstances := int(math.Ceil(float64(targetHeight) / float64(bgHeight)))
	for i := 0; i < yInstances; i++ {
		for j := 0; j < xInstances; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(j*bgWidth), float64(i*bgHeight))
			target.DrawImage(bg, op)
		}
	}
	return nil
}

func NewBackground(img string) Object {
	return &background{
		name: img,
	}
}
