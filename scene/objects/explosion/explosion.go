package explosion

import (
	"github.com/shamanr/battle_citty/interfaces"
)

type Explosion struct {
	interfaces.SceneObject
	counter int
}

func NewExplosion(obj interfaces.SceneObject) *Explosion {
	return &Explosion{SceneObject: obj}
}

func (e *Explosion) OnDraw() {
	e.counter++
	if e.counter % 5 == 0 {
		e.NextSprite()
	}

	if e.counter >= 5*3 {
		e.SceneObject.Delete()
	}
}
