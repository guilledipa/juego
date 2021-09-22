package assets

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"path"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Stall *Object
)

type Object struct {
	Image *ebiten.Image
	Specs *SpriteSheet
}
type SpriteSheet struct {
	Images []ImageSpec `json:"images"`
}
type ImageSpec struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	W    int    `json:"width"`
	H    int    `json:"height"`
}

func init() {
	Stall = &Object{
		Image: newImage(stallBytes),
		Specs: buildSpecs(relativePath("./stall.json")),
	}
}

func newImage(src []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(src))
	if err != nil {
		log.Fatalf("newImage: %v", err)
	}
	return ebiten.NewImageFromImage(img)
}

func buildSpecs(path string) *SpriteSheet {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("buildSpecs: %v", err)
	}
	s := &SpriteSheet{}
	json.Unmarshal(f, s)
	return s
}

func relativePath(filepath string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("relativePath(): error")
	}
	return path.Join(path.Dir(filename), filepath)
}
