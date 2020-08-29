package tank

import (
	"github.com/shamanr/battle_citty/consts"
	object "github.com/shamanr/battle_citty/scene/objects"
	"github.com/shamanr/battle_citty/scene/objects/projectile"
)


// NewTank возвращает объект танка
func NewTank(obj *object.MovableObject, bullet *projectile.Projectile) *Tank {
	return &Tank{
		MovableObject: obj,
		Bullet:        bullet,
	}
}

// Tank структруа танка
type Tank struct {
	*object.MovableObject
	Bullet *projectile.Projectile
}

func (t *Tank) Shoot() {
	t.Bullet.SetPos(t.GetPos())
	switch t.GetOrientation() {
	case consts.OrientationTop:
		t.Bullet.MoveUp()
		return
	case consts.OrientationRight:
		t.Bullet.MoveRight()
		return
	case consts.OrientationBottom:
		t.Bullet.MoveDown()
		return
	case consts.OrientationLeft:
		t.Bullet.MoveLeft()
		return
	}
}
