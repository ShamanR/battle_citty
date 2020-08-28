package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/interfaces"
	"github.com/shamanr/battle_citty/scene"
	"github.com/shamanr/battle_citty/scene/objects/tank"
)

type Game struct {
	scene  scene.Scene
	rm     interfaces.ResourceManager
	window *pixelgl.Window
}

func (g *Game) Init() {
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
	// Ищем точки РЕСПА ИГРОКА и Врагов
	var userSpawn interfaces.SceneObject
	var enemySpawns []interfaces.SceneObject
	userSpawn = nil
	for _, obj := range mapObjects {
		if obj.GetObjectType() == interfaces.ObjectTypePlayerSpawn {
			userSpawn = obj
			continue
		}
		if obj.GetObjectType() == interfaces.ObjectTypeAISpawn {
			enemySpawns = append(enemySpawns, obj)
		}
	}
	if userSpawn == nil {
		panic("userSpawn not found on map")
	}

	//
}

func (g *Game) MakeTank() *tank.Tank {
	obj := g.scene.MakeEmptyObj()
	obj.SetSpriteList(g.rm.GetSpriteMap(interfaces.ObjectTypePlayerTank1))
	return &tank.Tank{obj}
}

func (g *Game) MakeScene() interfaces.Scene {
	return nil
}

func (g *Game) MakeResourceManager() interfaces.ResourceManager {
	return nil
}
