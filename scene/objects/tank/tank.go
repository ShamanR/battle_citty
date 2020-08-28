package tank

import (
	"github.com/faiface/pixel"
	"math"
)

func NewTank(sprite *pixel.Sprite, pos pixel.Vec) *Tank {
	return &Tank{
		sprite: sprite,
		pos:    pos,
		angle:  0,
	}
}

type Tank struct {
	sprite *pixel.Sprite
	angle  float64
	pos    pixel.Vec
}

func (t *Tank) MoveLeft() {
	t.pos.X = t.pos.X - 5
	t.angle = math.Pi / 2
}

func (t *Tank) MoveRight() {
	t.pos.X = t.pos.X + 5
	t.angle = -math.Pi / 2
}

func (t *Tank) Draw(target pixel.Target) {
	//t.sprite.Frame().Moved(t.pos)
	mat := pixel.IM.Rotated(pixel.ZV, t.angle).Scaled(pixel.ZV, 3).Moved(t.pos)

	t.sprite.Draw(target, mat)
}
