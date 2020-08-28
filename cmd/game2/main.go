package main

import "github.com/shamanr/battle_citty/game"

func main() {
	game := game.Game{}
	game.Init()
	game.StartLevel()
}
