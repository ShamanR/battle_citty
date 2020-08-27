package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/actor"
	"github.com/shamanr/battle_citty/resource_loader"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"
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
	cfg := pixelgl.WindowConfig{
		Title:  "Battle City 1990",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	resourceLoader := resource_loader.Loader{
		ResourcePath: "./resources/",
	}
	resourceLoader.Init()

	tank := resourceLoader.MakeUserTank()
	tank.SetPosition(win.Bounds().Center())

	userActor := actor.User{
		Tank: tank,
	}
	go userActor.AttachToKeyboard(win)
	for !win.Closed() {
		win.Clear(colornames.Black)
		tank.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
