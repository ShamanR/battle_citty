package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/actors"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
	"github.com/shamanr/battle_citty/physics"
	"github.com/shamanr/battle_citty/resource_manager"
	"github.com/shamanr/battle_citty/scene"
	"github.com/shamanr/battle_citty/scene/objects/tank"
	"time"
)

type Game struct {
	scene   interfaces.Scene
	rm      interfaces.ResourceManager
	physics interfaces.Physics
	window  *pixelgl.Window
	player  actors.User
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
	g.window = win
	// Создаем СЦЕНУ
	g.scene = scene.NewScene()
	// Создаем ресурс-менеджер
	g.rm = resource_manager.NewResourceManager("resources/textures.png")
	// создаем физику
	g.physics = physics.New(time.Millisecond * 33)
	// ЗАГРУЖАЕМ НА СЦЕНУ КАРТУ
	//mapObjects := g.rm.LoadMap("")
	//for _, obj := range mapObjects {
	//	g.scene.AddObject(obj)
	//}
	//// Ищем точки РЕСПА ИГРОКА и Врагов
	//var userSpawn interfaces.SceneObject
	//var enemySpawns []interfaces.SceneObject
	//userSpawn = nil
	//for _, obj := range mapObjects {
	//	if obj.GetObjectType() == consts.ObjectTypePlayerSpawn {
	//		userSpawn = obj
	//		continue
	//	}
	//	if obj.GetObjectType() == consts.ObjectTypeAISpawn {
	//		enemySpawns = append(enemySpawns, obj)
	//	}
	//}
	userSpawn := g.scene.MakeEmptyObj()
	pos := g.window.Bounds().Center()
	userSpawn.SetPos(&pos)
	if userSpawn == nil {
		panic("userSpawn not found on map")
	}

	// Создаем объект танка
	playerTank := g.MakeTank()
	// Инстанцируем объект на сцену в точку респа
	playerTank.SetPos(userSpawn.GetPos())
	// PlayerActor
	player := actors.User{}
	player.SetTank(playerTank)
	g.player = player
}

func (g *Game) StartLevel() {
	last := time.Now()
	for !g.window.Closed() {
		dt := time.Since(last)
		<-time.After(time.Millisecond * 30)
		g.physics.MoveObjects(g.scene.GetObjects(), dt)
		g.scene.Draw(g.window)
		g.window.Update()
		last = time.Now()
	}
}

func (g *Game) MakeTank() *tank.Tank {
	obj := g.scene.MakeEmptyObj()
	obj.SetSpriteList(g.rm.GetSpriteMap(consts.ObjectTypePlayerTank1))
	return &tank.Tank{obj}
}
