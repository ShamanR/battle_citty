package interfaces

import (
	"time"

	"github.com/shamanr/battle_citty/consts"

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
}

type Scene interface {
	GetObjects() []SceneObject
	GetObjectByID(id int64) SceneObject
	SetSceneObjects(objects []SceneObject)
	GetSceneMap() SceneMap
	GetLevelMap() consts.LevelMap
	SetLevelMap(levelMap consts.LevelMap)
	MakeEmptyObj(objType consts.ObjectType) SceneObject
	Draw(target pixel.Target)
}
type Physics interface {
	MoveObjects(sceneMap SceneMap, dt time.Duration)
	PathTo(from, to pixel.Vec, sceneMap consts.LevelMap) []*pixel.Vec
}
