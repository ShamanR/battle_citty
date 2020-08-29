package brickwall

import (
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
)

// NewBrickWall возвращает объект стены
func NewBrickWall(obj interfaces.SceneObject) *BrickWall {
	bw := &BrickWall{obj}
	obj.SetGameObject(bw)
	return bw
}

// BrickWall структруа танка
type BrickWall struct {
	interfaces.SceneObject
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

func (t *BrickWall) OnDamage(other interfaces.SceneObject) {
	if other.GetObjectType() != consts.ObjectTypeProjectile {
		return
	}
	t.GetScene().RemoveObject(t.GetID())
	t.SceneObject.GetScene().GetSpawner().Spawn(consts.ObjectTypeExplosion, *t.SceneObject.GetPos())
}
