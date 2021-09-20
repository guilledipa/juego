package objects

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math"

	"github.com/guilledipa/juego/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type desk struct {
	name string
}

func (d *desk) Draw(target *ebiten.Image) error {
	img, format, err := image.Decode(bytes.NewReader(assets.AssetBytes(d.name)))
	if err != nil {
		log.Printf("deskDraw(\"%s\"): %s", d.name, format)
		return err
	}
	deskImg := ebiten.NewImageFromImage(img)

	targetWidth, targetHeight := target.Size()
	deskWidth, deskHeight := deskImg.Size()
	xInstances := int(math.Ceil(float64(targetWidth) / float64(deskWidth)))
	yPosition := targetHeight - deskHeight

	for j := 0; j < xInstances; j++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(j*deskWidth), float64(yPosition))
		target.DrawImage(deskImg, op)
	}
	return nil
}

func NewDesk(img string) Object {
	return &desk{
		name: img,
	}
}
