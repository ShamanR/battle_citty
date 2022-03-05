package game

import (
	actors2 "github.com/shamanr/battle_citty/internal/actors"
	consts2 "github.com/shamanr/battle_citty/internal/consts"
	"github.com/shamanr/battle_citty/internal/game/objects"
	"github.com/shamanr/battle_citty/internal/interfaces"
	"github.com/shamanr/battle_citty/internal/physics"
	"github.com/shamanr/battle_citty/internal/resource_manager"
	"github.com/shamanr/battle_citty/internal/scene"
	"github.com/shamanr/battle_citty/internal/scene/objects/brickwall"
	"github.com/shamanr/battle_citty/internal/scene/objects/explosion"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	scene              interfaces.Scene
	rm                 interfaces.ResourceManager
	physics            interfaces.Physics
	window             *pixelgl.Window
	player             *actors2.User
	ai                 []*actors2.AI
	lastID             int64
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

	g.rm.PlaySound(consts2.SoundGameIntro)
	// Ищем точки РЕСПА ИГРОКА и Врагов
	var userSpawn interfaces.SceneObject
	enemySpawns := []interfaces.SceneObject{}
	userSpawn = nil
	for _, obj := range g.scene.GetObjects() {
		if obj.GetObjectType() == consts2.ObjectTypePlayerSpawn {
			userSpawn = obj
			continue
		}
		if obj.GetObjectType() == consts2.ObjectTypeAISpawn {
			enemySpawns = append(enemySpawns, obj)
		}
	}

	// Создаем объект танка
	playerTank := g.gameObjectsManager.MakeTank(consts2.ObjectTypePlayerTank1)
	// Инстанцируем объект на сцену в точку респа
	playerTank.SetPos(userSpawn.GetPos())
	playerTank.SetScale(g.getScale())
	// PlayerActor
	player := actors2.User{}
	player.SetTank(playerTank)
	g.player = &player

	// AI
	// Инстанцируем объект на сцену в точку респа
	if enemySpawns != nil {
		for _, enemySpawn := range enemySpawns {
			ai := actors2.NewAI()
			aiTank := g.gameObjectsManager.MakeTank(consts2.ObjectTypePlayerTank1)
			enemyPos := enemySpawn.GetPos() //.Add(pixel.V(100, 140))
			aiTank.SetPos(enemyPos)
			aiTank.SetScale(g.getScale())
			ai.SetTank(aiTank)
			enemySpawn.SetVisible(false)

			g.ai = append(g.ai, ai)
		}
	}
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
		for _, ai := range g.ai {
			ai.Tick(dt)
		}
		g.physics.MoveObjects(g.scene.GetObjects(), dt)
		g.scene.Draw(g.window)
		g.window.Update()
	}
}

func (g *Game) getTileSize() pixel.Vec {
	tileSizeX := math.Round(g.window.Bounds().Max.X / consts2.MapSize)
	tileSizeY := math.Round(g.window.Bounds().Max.Y / consts2.MapSize)

	return pixel.V(tileSizeX, tileSizeY)
}

func (g *Game) getScale() pixel.Vec {
	tileSize := g.getTileSize()

	return pixel.V(tileSize.X/consts2.MapTileSize, tileSize.Y/consts2.MapTileSize)
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

func (g *Game) getGameObjectByType(typ consts2.ObjectType, pos pixel.Vec) interfaces.SceneObject {
	// Сейчас switch не очень нужен, но в будущем будем так создавать игровые объекты Танк, Стена и т.д.
	switch typ {
	case consts2.ObjectTypeBrickWall,
		consts2.ObjectTypeBrickWallDamagedLeft,
		consts2.ObjectTypeBrickWallDamagedTop,
		consts2.ObjectTypeBrickWallDamagedRight,
		consts2.ObjectTypeBrickWallDamagedDown:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetVisible(true)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		obj.SetGameObject(brickwall.NewBrickWall(obj))
		return obj
	case consts2.ObjectTypePlayerSpawn:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		return obj
	case consts2.ObjectTypeIronWall,
		consts2.ObjectTypeIronWallDamagedLeft,
		consts2.ObjectTypeIronWallDamagedTop,
		consts2.ObjectTypeIronWallDamagedRight,
		consts2.ObjectTypeIronWallDamagedDown:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetVisible(true)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		obj.SetGameObject(brickwall.NewBrickWall(obj))
		return obj
	case consts2.ObjectTypeHeadquarters:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetVisible(true)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		obj.SetGameObject(brickwall.NewBrickWall(obj))
		return obj
	case consts2.ObjectTypeExplosion:
		obj := g.scene.MakeEmptyObj(typ)
		obj.SetPos(&pos)
		obj.SetVisible(true)
		obj.SetSpriteList(g.rm.GetSpriteMap(typ))
		obj.SetGameObject(explosion.NewExplosion(obj))
		return obj
	case consts2.ObjectTypeEmpty:
		return nil
	}

	obj := g.scene.MakeEmptyObj(typ)
	obj.SetPos(&pos)
	obj.SetVisible(true)
	obj.SetSpriteList(g.rm.GetSpriteMap(typ))
	return obj
	//panic(errors.New(fmt.Sprintf("Unable to create object type %d", typ)))
}

func (g *Game) Spawn(objType consts2.ObjectType, pos pixel.Vec) {
	sceneObj := g.getGameObjectByType(objType, pos)
	scale := g.getScale()
	if sceneObj != nil {
		sceneObj.SetScale(scale)
	}
}
