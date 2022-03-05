package brickwall

import (
	consts2 "github.com/shamanr/battle_citty/internal/consts"
	"github.com/shamanr/battle_citty/internal/interfaces"
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
	t.SetOrientation(consts2.OrientationLeft)
}

// BreakRight разрушение стены справа
func (t *BrickWall) BreakRight() {
	t.SetOrientation(consts2.OrientationRight)
}

// BreakUp разрушение стены сверху
func (t *BrickWall) BreakUp() {
	t.SetOrientation(consts2.OrientationTop)
}

// BreakDown разрушение стены снизу
func (t *BrickWall) BreakDown() {
	t.SetOrientation(consts2.OrientationBottom)
}

func (t *BrickWall) OnDamage(other interfaces.SceneObject) {
	if other.GetObjectType() != consts2.ObjectTypeProjectile {
		return
	}
	t.GetScene().RemoveObject(t.GetID())
	t.SceneObject.GetScene().GetSpawner().Spawn(consts2.ObjectTypeExplosion, *t.SceneObject.GetPos())
}
