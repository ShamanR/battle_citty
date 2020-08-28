package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/resource_manager"
	object "github.com/shamanr/battle_citty/scene/objects"
	"github.com/shamanr/battle_citty/scene/objects/tank"
	"golang.org/x/image/colornames"
	_ "image/png"
	"time"
)

func run() {
	manager := resource_manager.NewResourceManager("resources/textures.png")
	sprite := manager.GetSprite(resource_manager.SimpleTankOrangeUp)

	cfg := pixelgl.WindowConfig{
		Title:  "Platformer",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	tank := tank.NewTank(sprite, 3)
	last := time.Now()
	for !win.Closed() {

		if win.Pressed(pixelgl.KeyLeft) {
			tank.MoveLeft()
		}
		if win.Pressed(pixelgl.KeyRight) {
			tank.MoveRight()
		}
		if win.JustPressed(pixelgl.KeyUp) {
			//ctrl.Y = 1
		}
		win.Clear(colornames.Black)
		win.SetSmooth(true)
		time.Since(last).Seconds()
		//mat := pixel.IM.Moved(win.GetSize().Center()).Scaled(pixel.ZV, 0.5)
		//sprite.Draw(win, mat)
		tank.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
