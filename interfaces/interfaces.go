package interfaces

import (
	"github.com/shamanr/battle_citty/consts"
	"time"

	"github.com/faiface/pixel"
)

// SceneMap ...
type SceneMap []SceneObject

/*
{
	"maps": {
		"1": [
			[ 1, 2, 3, 4]
		]
	}
}
*/

type ResourceManager interface {
	GetSpriteMap(name consts.ObjectType) *SceneObjectAnimateList
	LoadMap(path string) consts.LevelMap
}

// SceneObjectAnimateList структуры анимации (для танков?)
type SceneObjectAnimateList map[consts.Orientation][]*pixel.Sprite

// SceneObject интерфейс
type SceneObject interface {
	GetID() int64
	GetPos() *pixel.Vec
	SetPos(vect *pixel.Vec)
	GetSpeed() *pixel.Vec
	SetSpeed(vect *pixel.Vec)
	GetSprite() *pixel.Sprite
	SetSpriteList(list *SceneObjectAnimateList)
	Draw(target pixel.Target)
	SetScale(scale float64)
	GetScale() float64
	GetSize() *pixel.Rect
	IsVisible() bool
	SetVisible(visible bool)
	GetObjects() []SceneObject
	GetObjectType() consts.ObjectType
	SetOrientation(orient consts.Orientation)
	GetOrientation() consts.Orientation
	Delete()
}

type Scene interface {
	GetObjects() []SceneObject
	GetObjectByID(id int64) SceneObject
	FillSceneObjectsByMap(levelMap consts.LevelMap)
	GetSceneMap() SceneMap
	GetLevelMap() consts.LevelMap
	MakeEmptyObj(objType consts.ObjectType) SceneObject
	Draw(target pixel.Target)
}
type Physics interface {
	MoveObjects(sceneMap SceneMap, dt time.Duration)
	PathTo(from, to pixel.Vec, sceneMap consts.LevelMap)
}
