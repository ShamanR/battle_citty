package tank

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/scene/objects"
)

const (
	OrientationTop   = 0
	OrientationLeft  = 1
	OrientationDown  = 2
	OrientationRight = 3
)

type Tank struct {
	objects.Object
	TankAnimation map[int][]*pixel.Sprite
	orientation   int8
}

func (t *Tank) MoveLeft() {
	if t.orientation != OrientationLeft {
		t.SetAnimation(t.TankAnimation[OrientationLeft])
		t.orientation = OrientationLeft
	}
	pos := t.GetPosition()
	pos.X -= 2
	t.SetPosition(pos)
	t.NextAnimationFrame()
}

func (t *Tank) MoveRight() {
	if t.orientation != OrientationRight {
		t.SetAnimation(t.TankAnimation[OrientationRight])
		t.orientation = OrientationRight
	}
	pos := t.GetPosition()
	pos.X += 2
	t.SetPosition(pos)
	t.NextAnimationFrame()
}

func (t *Tank) MoveTop() {
	if t.orientation != OrientationTop {
		t.SetAnimation(t.TankAnimation[OrientationTop])
		t.orientation = OrientationTop
	}
	pos := t.GetPosition()
	pos.Y += 2
	t.SetPosition(pos)
	t.NextAnimationFrame()
}

func (t *Tank) MoveDown() {
	if t.orientation != OrientationDown {
		t.SetAnimation(t.TankAnimation[OrientationDown])
		t.orientation = OrientationDown
	}
	pos := t.GetPosition()
	pos.Y -= 2
	t.SetPosition(pos)
	t.NextAnimationFrame()
}

func (t *Tank) Stop() {
}

func (t *Tank) Shoot() {
	// todo
}
