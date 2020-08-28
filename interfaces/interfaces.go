package interfaces

import (
	"github.com/faiface/pixel"
)

// ObjectType тип объекта
type ObjectType uint8

// SceneMap ...
type SceneMap []SceneObject

// LevelMap ...
type LevelMap [][]ObjectType

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
	OrientationTop    = 0
	OrientationRight  = 1
	OrientationBottom = 2
	OrientationLeft   = 3
)

type ResourceManager interface {
	GetSprite(name string) *pixel.Sprite
	LoadMap() SceneMap
}

// SceneObjectAnimateList структуры анимации (для танков?)
type SceneObjectAnimateList struct {
	LeftSprite   []*pixel.Sprite
	RightSprite  []*pixel.Sprite
	TopSprite    []*pixel.Sprite
	BottomSprite []*pixel.Sprite
}

// SceneObject интерфейс
type SceneObject interface {
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
	GetObjectById() SceneObject
	AddObject(object SceneObject)
	GetSceneMap() SceneMap
	GetLevelMap() LevelMap
	Draw(target pixel.Target)
}
type Physics interface {
	MoveObjects(sceneMap SceneMap)
	canCollide(obj, obj2 SceneObject) bool
	PathTo(from, to pixel.Vec, sceneMap LevelMap)
}
