package main

import (
	"log"
	"math/rand"
	"time"

	tiropato "github.com/guilledipa/juego"
)

func main() {
	rand.Seed(time.Now().Unix())
	game := tiropato.NewGame()
	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
