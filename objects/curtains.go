package objects

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/guilledipa/juego/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type curtains struct {
	name string
}

func (d *curtains) Draw(target *ebiten.Image) error {
	img, format, err := image.Decode(bytes.NewReader(assets.AssetBytes(d.name)))
	if err != nil {
		log.Printf("curtainsDraw(\"%s\"): %s", d.name, format)
		return err
	}
	curtainsImg := ebiten.NewImageFromImage(img)

	targetWidth, _ := target.Size()

	op := &ebiten.DrawImageOptions{}
	target.DrawImage(curtainsImg, op)
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(targetWidth), float64(0))
	target.DrawImage(curtainsImg, op)
	return nil
}

func NewCurtains(img string) Object {
	return &curtains{
		name: img,
	}
}
