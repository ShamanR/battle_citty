package projectile

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/interfaces"
	"github.com/shamanr/battle_citty/scene/objects"
)

// NewProjectile возвращает объект пули
func NewProjectile(obj *object.MovableObject) *Projectile {
	p := &Projectile{
		MovableObject: obj,
	}
	obj.SceneObject.SetGameObject(p)
	return p
}

// MovableObject структруа танка
type Projectile struct {
	*object.MovableObject
}

func (p *Projectile) OnDamage(other interfaces.SceneObject) {
	p.MovableObject.SetPos(&pixel.Vec{
		X: -100,
		Y: -110,
	})
	p.MovableObject.SetSpeed(&pixel.Vec{
		X: 0,
		Y: 0,
	})
}

