package tank

import (
	consts2 "github.com/shamanr/battle_citty/internal/consts"
	"github.com/shamanr/battle_citty/internal/interfaces"
	"github.com/shamanr/battle_citty/internal/scene/objects"
	"github.com/shamanr/battle_citty/internal/scene/objects/projectile"
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
	t.Bullet.SetScale(t.GetScale())
	t.Bullet.SetVisible(true)
	t.AddChild(t.Bullet)
	startPos := *t.GetPos()
	switch t.GetOrientation() {
	case consts2.OrientationTop:
		startPos.Y += t.GetSize().H()
		t.Bullet.SetPos(&startPos)
		t.Bullet.MoveUp()
		return
	case consts2.OrientationRight:
		startPos.X += t.GetSize().W()
		t.Bullet.SetPos(&startPos)
		t.Bullet.MoveRight()
		return
	case consts2.OrientationBottom:
		startPos.Y -= t.GetSize().H()
		t.Bullet.SetPos(&startPos)
		t.Bullet.MoveDown()
		return
	case consts2.OrientationLeft:
		startPos.X -= t.GetSize().W()
		t.Bullet.SetPos(&startPos)
		t.Bullet.MoveLeft()
		return
	}
}

func (t *Tank) OnDamage(other interfaces.SceneObject) {
	if other.GetObjectType() != consts2.ObjectTypeProjectile {
		return
	}

	childs := t.GetObjects()
	for _, ch := range childs {
		if ch.GetID() == other.GetID() {
			return
		}
	}

	t.SceneObject.SetLife(t.SceneObject.GetLife() - 1)
	if t.SceneObject.GetLife() <= 0 || t.SceneObject.GetLife() > 100 {
		t.SceneObject.GetScene().GetSpawner().Spawn(consts2.ObjectTypeExplosion, *t.SceneObject.GetPos())
		t.Delete()
	}
}
