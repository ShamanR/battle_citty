package game

import (
	"errors"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/actors"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
	"github.com/shamanr/battle_citty/physics"
	"github.com/shamanr/battle_citty/resource_manager"
	"github.com/shamanr/battle_citty/scene"
	object "github.com/shamanr/battle_citty/scene/objects"
	"github.com/shamanr/battle_citty/scene/objects/tank"
	"math"
	"time"
)

type Game struct {
	scene   interfaces.Scene
	rm      interfaces.ResourceManager
	physics interfaces.Physics
	window  *pixelgl.Window
	player  actors.User

	lastID int64
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
	g.lastID = 0
	// Создаем СЦЕНУ
	g.scene = scene.NewScene()
	// Создаем ресурс-менеджер
	g.rm = resource_manager.NewResourceManager("resources/textures.png")
	// создаем физику
	g.physics = physics.New(33*time.Millisecond, 16, 3)
	// Стартуем первый уровень
	g.fillSceneByMap("resources/level1.json")
	// Ищем точки РЕСПА ИГРОКА и Врагов
	var userSpawn interfaces.SceneObject
	var enemySpawns []interfaces.SceneObject
	userSpawn = nil
	for _, obj := range g.scene.GetObjects() {
		if obj.GetObjectType() == consts.ObjectTypePlayerSpawn {
			userSpawn = obj
			continue
		}
		if obj.GetObjectType() == consts.ObjectTypeAISpawn {
			enemySpawns = append(enemySpawns, obj)
		}
	}

	// Создаем объект танка
	playerTank := g.MakeTank()
	// Инстанцируем объект на сцену в точку респа
	playerTank.SetPos(userSpawn.GetPos())
	playerTank.SetScale(g.getScale())
	// PlayerActor
	player := actors.User{}
	player.SetTank(playerTank)
	g.player = player
}

func (g *Game) StartLevel() {
	last := time.Now()
	g.lastID = 0
	for !g.window.Closed() {
		dt := time.Since(last)
		last = time.Now()
		<-time.After(time.Millisecond * 30)
		g.player.AttachToKeyboard(g.window)
		g.physics.MoveObjects(g.scene.GetObjects(), dt)
		g.scene.Draw(g.window)
		g.window.Update()
	}
}

func (g *Game) getTileSize() pixel.Vec {
	tileSizeX := math.Round(g.window.Bounds().Max.X / consts.MapSize)
	tileSizeY := math.Round(g.window.Bounds().Max.Y / consts.MapSize)

	return pixel.V(tileSizeX, tileSizeY)
}

func (g *Game) getScale() pixel.Vec {
	tileSize := g.getTileSize()

	return pixel.V(tileSize.X / consts.MapTileSize, tileSize.Y / consts.MapTileSize)
}

func (g *Game) getNewId() int64 {
	g.lastID++
	return g.lastID
}

func (g *Game) fillSceneByMap(levelMapPath string) {
	levelMap := g.rm.LoadMap(levelMapPath)

	var sceneObjects []interfaces.SceneObject
	for y, row := range levelMap {
		for x, objType := range row {
			tileSize := g.getTileSize()
			currentPos := pixel.V(float64(x) * tileSize.X, g.window.Bounds().Max.Y - float64(y) * tileSize.Y)

			sceneObj := g.getGameObjectByType(objType, currentPos)
			scale := g.getScale()
			if sceneObj != nil {
				// Увеличиваем спрайт в соответствии с размером игрового поля
				sceneObj.SetScale(scale)
				sceneObjects = append(sceneObjects, sceneObj)
			}
		}
	}

	g.scene.SetSceneObjects(sceneObjects)
}

func (g *Game) getGameObjectByType(typ consts.ObjectType, pos pixel.Vec) interfaces.SceneObject {
	// Сейчас switch не очень нужен, но в будущем будем так создавать игровые объекты Танк, Стена и т.д.
	switch typ {
	case consts.ObjectTypeBrickWall:
		return object.NewObject(g.getNewId(), typ, &pos, g.rm.GetSpriteMap(typ))
	case consts.ObjectTypePlayerSpawn:
		return object.NewObject(g.getNewId(), typ, &pos, g.rm.GetSpriteMap(typ))
	case consts.ObjectTypeEmpty:
		return nil
	}

	panic(errors.New(fmt.Sprintf("Unable to create object type %s", typ)))
}

func (g *Game) MakeTank() *tank.Tank {
	obj := g.scene.MakeEmptyObj(consts.ObjectTypePlayerTank1)
	obj.SetSpriteList(g.rm.GetSpriteMap(consts.ObjectTypePlayerTank1))
	return tank.NewTank(obj, 5)
}
