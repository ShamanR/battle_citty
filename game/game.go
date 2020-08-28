package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/interfaces"
	"github.com/shamanr/battle_citty/resource_manager"
)

type Game struct {
}

func (g *Game) Run() {
	// СОЗДАЕМ ОКНО ИГРЫ
	cfg := pixelgl.WindowConfig{
		Title:  "Platformer",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	// Создаем СЦЕНУ и ресурс-менеджер
	scene := g.MakeScene()
	rm := g.MakeResourceManager()
	// ЗАГРУЖАЕМ НА СЦЕНУ КАРТУ
	mapObjects := rm.LoadMap()
	for _, obj := range mapObjects {
		scene.AddObject(obj)
	}
	// Ищем точки РЕСПА
	var userResp interfaces.SceneObject
	for _, obj := range mapObjects {
		if obj.GetObjectType() == resource_manager.Spawn {
			userResp = obj
		}
	}
}

func (g *Game) MakeScene() interfaces.Scene {
	return nil
}

func (g *Game) MakeResourceManager() interfaces.ResourceManager {
	return nil
}
