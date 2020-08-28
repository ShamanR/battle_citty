package interfaces

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/scene/objects/tank"
)

type ObjectType uint8
type SceneMap []SceneObject
type LevelMap [][]ObjectType
type SpriteType string
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
	BrickWall ObjectType = iota
	IronWall
	Water
	Forest
	Ice
	Headquarters
	PlayerSpawn
	AISpawn
	Bonus
	SimpleOrangeTank
)

type ResourceManager interface {
	GetSprite(name SpriteType) *pixel.Sprite
	GetSpriteMap(name ObjectType) *SceneObjectAnimateList
	LoadMap() *SceneMap
	MakeTank(name ObjectType) *tank.Tank
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
