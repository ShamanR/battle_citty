package interfaces

import (
	"time"

	"github.com/faiface/pixel"
)

// ObjectType тип объекта
type ObjectType uint8

// SceneMap ...
type SceneMap []SceneObject

// LevelMap ...
type LevelMap [][]ObjectType
type SpriteType string

// Orientation ориентация
type Orientation uint8

/*
{
	"maps": {
		"1": [
			[ 1, 2, 3, 4]
		]
	}
}
*/

const (
	ObjectTypeBrickWall ObjectType = iota
	ObjectTypeIronWall
	ObjectTypeWater
	ObjectTypeForest
	ObjectTypeIce
	ObjectTypeHeadquarters
	ObjectTypePlayerSpawn
	ObjectTypeAISpawn
	ObjectTypeBonus
	ObjectTypePlayerTank1
	ObjectTypePlayerTank2
	ObjectTypePlayerTank3
	ObjectTypePlayerTank4
	ObjectTypeEnemyTank1
	ObjectTypeEnemyTank2
	ObjectTypeEnemyTank3
	ObjectTypeEnemyTank4
	ObjectTypeProjectile
)

const (
	OrientationTop    = 0
	OrientationRight  = 1
	OrientationBottom = 2
	OrientationLeft   = 3
)

type ResourceManager interface {
	GetSprite(name SpriteType) *pixel.Sprite
	GetSpriteMap(name ObjectType) *SceneObjectAnimateList
	LoadMap() *SceneMap
	MakeTank(name ObjectType) SceneObject
}

// SceneObjectAnimateList структуры анимации (для танков?)
type SceneObjectAnimateList map[Orientation][]*pixel.Sprite

// SceneObject интерфейс
type SceneObject interface {
	GetId() int64
	GetPos() *pixel.Vec
	SetPos(vect *pixel.Vec)
	GetSpeed() *pixel.Vec
	SetSpeed(vect *pixel.Vec)
	getSprite() *pixel.Sprite
	SetSpriteList(list *SceneObjectAnimateList)
	getSriteListLen() int64
	Draw(target pixel.Target)
	SetScale(scale float64)
	GetScale() float64
	Bounds() *pixel.Rect
	IsVisible() bool
	SetVisible(visible bool)
	GetObjects() []SceneObject
	GetObjectType() ObjectType
	SetOrientation(orient Orientation)
	Delete()
}

type Scene interface {
	GetObjects() []SceneObject
	GetObjectByID(id int64) SceneObject
	AddObject(object SceneObject)
	GetSceneMap() SceneMap
	GetLevelMap() LevelMap
	Draw(target pixel.Target)
}
type Physics interface {
	MoveObjects(sceneMap SceneMap, dt time.Duration)
	canCollide(obj, obj2 SceneObject) bool
	PathTo(from, to pixel.Vec, sceneMap LevelMap)
}
