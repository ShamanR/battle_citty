package tank

import (
	object "github.com/shamanr/battle_citty/scene/objects"

	"github.com/faiface/pixel"
)

func NewTank() *Tank {
	return &Tank{}
}

type Tank struct {
	object.Object
}

// MoveLeft передвигает танк влево
func (t *Tank) MoveLeft() {
	v := t.GetPos()
	vn := pixel.V(v.X-5, v.Y)
	t.SetPos(&vn)
	// t.pos.X = t.pos.X - 5
	t.SetOrientation(object.OrientationLeft)
}

// MoveRight передвигает танк вправо
func (t *Tank) MoveRight() {
	v := t.GetPos()
	vn := pixel.V(v.X+5, v.Y)
	t.SetPos(&vn)
	t.SetOrientation(object.OrientationRight)
}

// MoveUp передвигает танк влево
func (t *Tank) MoveUp() {
	v := t.GetPos()
	vn := pixel.V(v.X, v.Y+5)
	t.SetPos(&vn)
	// t.pos.X = t.pos.X - 5
	t.SetOrientation(object.OrientationTop)
}

// MoveDown передвигает танк вправо
func (t *Tank) MoveDown() {
	v := t.GetPos()
	vn := pixel.V(v.X, v.Y-5)
	t.SetPos(&vn)
	t.SetOrientation(object.OrientationBottom)
}
