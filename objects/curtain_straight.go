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

type curtainStraight struct {
	name string
}

func (c *curtainStraight) Draw(target *ebiten.Image) error {
	img, format, err := image.Decode(bytes.NewReader(assets.AssetBytes(c.name)))
	if err != nil {
		log.Printf("curtainStraightDraw(\"%s\"): %s", c.name, format)
		return err
	}
	curtainStraightImg := ebiten.NewImageFromImage(img)

	targetWidth, _ := target.Size()
	curtainStraightWidth, _ := curtainStraightImg.Size()
	xInstances := int(math.Ceil(float64(targetWidth) / float64(curtainStraightWidth)))

	for i := 0; i < xInstances; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i*curtainStraightWidth), float64(0))
		target.DrawImage(curtainStraightImg, op)
	}
	return nil
}

func NewCurtainStraight(img string) Object {
	return &curtainStraight{
		name: img,
	}
}
