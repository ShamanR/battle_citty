package interfaces

import (
	"time"

	"github.com/shamanr/battle_citty/consts"

	"github.com/faiface/pixel"
)

// SceneMap ...
type SceneMap []SceneObject

type ResourceManager interface {
	GetSpriteMap(name consts.ObjectType) *SceneObjectAnimateList
	LoadMap(path string) consts.LevelMap
	PlaySound(name consts.SoundType)
	CloseSound()
}

// SceneObjectAnimateList структуры анимации (для танков?)
type SceneObjectAnimateList map[consts.Orientation][]*pixel.Sprite

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
	GetObjects() []SceneObject
	GetObjectType() consts.ObjectType
	SetOrientation(orient consts.Orientation)
	GetOrientation() consts.Orientation
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
	GetLevelMap() consts.LevelMap
	SetLevelMap(levelMap consts.LevelMap)
	MakeEmptyObj(objType consts.ObjectType) SceneObject
	Draw(target pixel.Target)
	GetSpawner() Spawner
}
type Physics interface {
	MoveObjects(gameMap SceneMap, dt time.Duration)
	PathTo(from, to pixel.Vec, sceneMap consts.LevelMap) []*pixel.Vec
}

type Damageable interface {
	OnDamage(other SceneObject)
}

type Spawner interface {
	Spawn(objType consts.ObjectType, pos pixel.Vec)
}
