package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/resources"
	"github.com/shamanr/battle_citty/scene/objects/tank"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"
	"time"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open("resources/" + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	// ресурсы
	// open and load the spritesheet
	manager := resources.NewResourceManager("resources/textures.png")
	sprite := manager.GetSprite(resources.SimpleTankOrangeUp)

	cfg := pixelgl.WindowConfig{
		Title:  "Platformer",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	tank := tank.NewTank(sprite, win.Bounds().Center())
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
		//mat := pixel.IM.Moved(win.Bounds().Center()).Scaled(pixel.ZV, 0.5)
		//sprite.Draw(win, mat)
		tank.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
