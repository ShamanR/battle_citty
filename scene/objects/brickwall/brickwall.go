package brickwall

import (
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
	object "github.com/shamanr/battle_citty/scene/objects"

	"github.com/faiface/pixel"
)

// NewBrickWall возвращает объект стены
func NewBrickWall(Id int64, scene interfaces.Scene, objectType consts.ObjectType, pos *pixel.Vec, spriteList *interfaces.SceneObjectAnimateList) *BrickWall {
	obj := object.NewObject(Id, scene, objectType, pos, spriteList)
	return &BrickWall{obj}
}

// BrickWall структруа танка
type BrickWall struct {
	*object.Object
}

// BreakLeft разрушение стены слева
func (t *BrickWall) BreakLeft() {
	t.SetOrientation(consts.OrientationLeft)
}

// BreakRight разрушение стены справа
func (t *BrickWall) BreakRight() {
	t.SetOrientation(consts.OrientationRight)
}

// BreakUp разрушение стены сверху
func (t *BrickWall) BreakUp() {
	t.SetOrientation(consts.OrientationTop)
}

// BreakDown разрушение стены снизу
func (t *BrickWall) BreakDown() {
	t.SetOrientation(consts.OrientationBottom)
}
