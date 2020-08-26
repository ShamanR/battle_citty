package tank

import (
	"github.com/shamanr/battle_citty/scene/objects"
	"math"
)

type UserTank struct {
	objects.Object
}

func (t *UserTank) MoveLeft() {
	t.SetAngle(math.Pi / 2)
	pos := t.GetPosition()
	pos.X -= 2
	t.SetPosition(pos)
	t.NextAnimationFrame()
}

func (t *UserTank) MoveRight() {
	t.SetAngle(-math.Pi / 2)
	pos := t.GetPosition()
	pos.X += 2
	t.SetPosition(pos)
	t.NextAnimationFrame()
}

func (t *UserTank) MoveTop() {
	t.SetAngle(0)
	pos := t.GetPosition()
	pos.Y += 2
	t.SetPosition(pos)
	t.NextAnimationFrame()
}

func (t *UserTank) MoveDown() {
	t.SetAngle(math.Pi)
	pos := t.GetPosition()
	pos.Y -= 2
	t.SetPosition(pos)
	t.NextAnimationFrame()
}

func (t *UserTank) Shoot() {
	// todo
}
