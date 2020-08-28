package interfaces

import (
	"github.com/faiface/pixel"
)

type ObjectType uint8
type SceneMap []SceneObject
type LevelMap [][]ObjectType

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
	GetSprite(name string) *pixel.Sprite
	LoadMap() SceneMap
}
type SceneObjectAnimateList struct {
	LeftSprite   []*pixel.Sprite
	RightSprite  []*pixel.Sprite
	TopSprite    []*pixel.Sprite
	BottomSprite []*pixel.Sprite
}
type SceneObject interface {
	GetID() int64
	GetPos() *pixel.Vec
	SetPos(vect *pixel.Vec)
	GetSpeed() *pixel.Vec
	SetSpeed(vect *pixel.Vec)
	getSprite() *pixel.Sprite
	SetSpriteList(list *SceneObjectAnimateList)
	Draw(target pixel.Target)
	SetScale(scale float64)
	GetScale(scale float64)
	Bounds() pixel.Rect
	IsVisible() bool
	SetVisible(visible bool)
	GetObjects() []SceneObject
	GetObjectType() ObjectType
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
