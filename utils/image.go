package utils

import (
	"fmt"
	"image"

	"github.com/guilledipa/juego/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func GetImage(name string, obj *assets.Object) (*ebiten.Image, error) {
	var rect image.Rectangle
	for _, img := range obj.Specs.Images {
		if img.Name == name {
			rect = image.Rect(img.X, img.Y, img.X+img.W, img.Y+img.H)
			break
		} else {
			return nil, fmt.Errorf("GetImage(\"%s\": Unable to find image in Spec", name)
		}
	}
	return obj.Image.SubImage(rect).(*ebiten.Image), nil
}
