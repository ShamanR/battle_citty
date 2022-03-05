package interfaces

import (
	consts2 "github.com/shamanr/battle_citty/internal/consts"
	"time"

	"github.com/faiface/pixel"
)

// SceneMap ...
type SceneMap []SceneObject

type ResourceManager interface {
	GetSpriteMap(name consts2.ObjectType) *SceneObjectAnimateList
	LoadMap(path string) consts2.LevelMap
	PlaySound(name consts2.SoundType)
	CloseSound()
}

// SceneObjectAnimateList структуры анимации (для танков?)
type SceneObjectAnimateList map[consts2.Orientation][]*pixel.Sprite

// SceneObject интерфейс
type SceneObject interface {
	GetID() int64
	GetScene() Scene
	GetPos() *pixel.Vec
	SetPos(vect *pixel.Vec)
	GetSpeed() *pixel.Vec
	SetSpeed(vect *pixel.Vec)
	GetSprite() *pixel.Sprite
	SetSpriteList(list *SceneObjectAnimateList)
	Draw(target pixel.Target)
	SetScale(scale pixel.Vec)
	GetScale() pixel.Vec
	GetSize() *pixel.Rect
	IsVisible() bool
	SetVisible(visible bool)
	AddChild(obj SceneObject)
	GetObjects() []SceneObject
	GetObjectType() consts2.ObjectType
	SetOrientation(orient consts2.Orientation)
	GetOrientation() consts2.Orientation
	NextSprite()
	Delete()
	OnCollide(obj SceneObject) // Коллизия произошла, вызываем этот метод
	GetGameObject() interface{}
	SetGameObject(gObj interface{})
	SetLife(life uint8)
	GetLife() uint8
}

type Scene interface {
	GetObjects() []SceneObject
	GetObjectByID(id int64) SceneObject
	RemoveObject(id int64)
	SetSceneObjects(objects []SceneObject)
	GetSceneMap() SceneMap
	GetLevelMap() consts2.LevelMap
	SetLevelMap(levelMap consts2.LevelMap)
	MakeEmptyObj(objType consts2.ObjectType) SceneObject
	Draw(target pixel.Target)
	GetSpawner() Spawner
}
type Physics interface {
	MoveObjects(gameMap SceneMap, dt time.Duration)
	PathTo(from, to pixel.Vec, sceneMap consts2.LevelMap) []*pixel.Vec
}

type Damageable interface {
	OnDamage(other SceneObject)
}

type Spawner interface {
	Spawn(objType consts2.ObjectType, pos pixel.Vec)
}

type DrawListener interface {
	OnDraw()
}
