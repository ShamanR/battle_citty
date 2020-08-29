package game

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/shamanr/battle_citty/actors"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/game/objects"
	"github.com/shamanr/battle_citty/interfaces"
	"github.com/shamanr/battle_citty/physics"
	"github.com/shamanr/battle_citty/resource_manager"
	"github.com/shamanr/battle_citty/scene"
	"golang.org/x/image/colornames"
)

type Game struct {
	scene   interfaces.Scene
	rm      interfaces.ResourceManager
	physics interfaces.Physics
	window  *pixelgl.Window
	player  *actors.User
	ai      *actors.AI
	lastID  int64
	gameObjectsManager *objects.GameObjectsManager
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
	g.scene = scene.NewScene(g)
	// Создаем ресурс-менеджер
	g.rm = resource_manager.NewResourceManager("resources/textures.png")
	// создаем физику
	g.physics = physics.New(33*time.Millisecond, 16, 3)
	// Стартуем первый уровень
	g.fillSceneByMap("resources/level1.json")
	// Создаем менеджер игровых объектов
	g.gameObjectsManager = objects.NewGameObjectsManager(g.rm, g.scene)
	g.rm.PlaySound(consts.SoundGameIntro)
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
	playerTank := g.gameObjectsManager.MakeTank(consts.ObjectTypePlayerTank1)
	// Инстанцируем объект на сцену в точку респа
	playerTank.SetPos(userSpawn.GetPos())
	playerTank.SetScale(g.getScale())
	// PlayerActor
	player := actors.User{}
	player.SetTank(playerTank)
	g.player = &player

	// AI
	g.ai = actors.NewAI()
	aiTank := g.gameObjectsManager.MakeTank(consts.ObjectTypePlayerTank1)
	// Инстанцируем объект на сцену в точку респа
	enemyPos := userSpawn.GetPos().Add(pixel.V(100, 140))
	aiTank.SetPos(&enemyPos)
	aiTank.SetScale(g.getScale())
	g.ai.SetTank(aiTank)
}

func (g *Game) StartLevel() {
	last := time.Now()
	defer g.rm.CloseSound()
	for !g.window.Closed() {
		dt := time.Since(last)
		last = time.Now()
		<-time.After(time.Millisecond * 30)
		g.window.Clear(colornames.Black)
		g.player.AttachToKeyboard(g.window)
		g.ai.Tick(dt)
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

	return pixel.V(tileSize.X/consts.MapTileSize, tileSize.Y/consts.MapTileSize)
}

func (g *Game) fillSceneByMap(levelMapPath string) {
	levelMap := g.rm.LoadMap(levelMapPath)

	var sceneObjects []interfaces.SceneObject
	for y, row := range levelMap {
		for x, objType := range row {
			tileSize := g.getTileSize()
			currentPos := pixel.V(float64(x)*tileSize.X+tileSize.X/2, g.window.Bounds().Max.Y-float64(y)*tileSize.Y-tileSize.Y/2)

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
	g.scene.SetLevelMap(levelMap)
}

func (g *Game) getGameObjectByType(typ consts.ObjectType, pos pixel.Vec) interfaces.SceneObject {
	// Сейчас switch не очень нужен, но в будущем будем так создавать игровые объекты Танк, Стена и т.д.
	switch typ {
	case consts.ObjectTypeBrickWall:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetVisible(true)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		return obj
	case consts.ObjectTypePlayerSpawn:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		return obj
	case consts.ObjectTypeIronWall:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetVisible(true)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		return obj
	case consts.ObjectTypeHeadquarters:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetVisible(true)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		return obj
	case consts.ObjectTypeEmpty:
		return nil
	}

	obj := g.scene.MakeEmptyObj(typ)
	obj.SetPos(&pos)
	obj.SetVisible(true)
	obj.SetSpriteList(g.rm.GetSpriteMap(typ))
	return obj
	//panic(errors.New(fmt.Sprintf("Unable to create object type %d", typ)))
}

func (g *Game) Spawn(objType consts.ObjectType, pos pixel.Vec) {
	sceneObj := g.getGameObjectByType(objType, pos)
	scale := g.getScale()
	if sceneObj != nil {
		sceneObj.SetScale(scale)
	}
}

