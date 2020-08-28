package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/game"
)

func run() {
	game := game.Game{}
	game.Init()
	game.StartLevel()
}

func main() {
	pixelgl.Run(run)
}
