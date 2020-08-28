package brickwall

import (
	"github.com/shamanr/battle_citty/interfaces"
	object "github.com/shamanr/battle_citty/scene/objects"

	"github.com/faiface/pixel"
)

// NewBrickWall возвращает объект стены
func NewBrickWall(objectType interfaces.ObjectType, pos *pixel.Vec, spriteList *interfaces.SceneObjectAnimateList) *BrickWall {
	obj := object.NewObject(objectType, pos, spriteList)
	return &BrickWall{obj}
}

// BrickWall структруа танка
type BrickWall struct {
	*object.Object
}

// BreakLeft разрушение стены слева
func (t *BrickWall) BreakLeft() {
	t.SetOrientation(interfaces.OrientationLeft)
}

// BreakRight разрушение стены справа
func (t *BrickWall) BreakRight() {
	t.SetOrientation(interfaces.OrientationRight)
}

// BreakUp разрушение стены сверху
func (t *BrickWall) BreakUp() {
	t.SetOrientation(interfaces.OrientationTop)
}

// BreakDown разрушение стены снизу
func (t *BrickWall) BreakDown() {
	t.SetOrientation(interfaces.OrientationBottom)
}
