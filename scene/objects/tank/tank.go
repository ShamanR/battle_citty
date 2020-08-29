package tank

import (
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
	object "github.com/shamanr/battle_citty/scene/objects"
	"github.com/shamanr/battle_citty/scene/objects/projectile"
)


// NewTank возвращает объект танка
func NewTank(obj *object.MovableObject, bullet *projectile.Projectile) *Tank {
	t := &Tank{
		MovableObject: obj,
		Bullet:        bullet,
	}
	obj.SceneObject.SetGameObject(t)
	return t
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

func (t *Tank) OnDamage(other interfaces.SceneObject) {
	t.SceneObject.SetLife(t.SceneObject.GetLife() - 1)
	if t.SceneObject.GetLife() <= 0 {
		t.SceneObject.GetScene().GetSpawner().Spawn(consts.ObjectTypeExplosion, *t.SceneObject.GetPos())
		t.Delete()
	}
}
