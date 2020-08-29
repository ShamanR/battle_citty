package projectile

import (
	"github.com/shamanr/battle_citty/scene/objects"
)

// NewProjectile возвращает объект пули
func NewProjectile(obj *object.MovableObject) *Projectile {
	return &Projectile{
		MovableObject: obj,
	}
}

// MovableObject структруа танка
type Projectile struct {
	*object.MovableObject
}

