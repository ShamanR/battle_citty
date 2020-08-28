package tank

import (
	"github.com/shamanr/battle_citty/interfaces"
	object "github.com/shamanr/battle_citty/scene/objects"

	"github.com/faiface/pixel"
)

// NewTank возвращает объект танка
func NewTank(objectType interfaces.ObjectType, pos *pixel.Vec, spriteList *interfaces.SceneObjectAnimateList) *Tank {
	obj := object.NewObject(objectType, pos, spriteList)
	return &Tank{obj}
}

// Tank структруа танка
type Tank struct {
	*object.Object
}

// MoveLeft передвигает танк влево
func (t *Tank) MoveLeft() {
	v := t.GetPos()
	vn := pixel.V(v.X-5, v.Y)
	t.SetPos(&vn)
	// t.pos.X = t.pos.X - 5
	t.SetOrientation(interfaces.OrientationLeft)
}

// MoveRight передвигает танк вправо
func (t *Tank) MoveRight() {
	v := t.GetPos()
	vn := pixel.V(v.X+5, v.Y)
	t.SetPos(&vn)
	t.SetOrientation(interfaces.OrientationRight)
}

// MoveUp передвигает танк влево
func (t *Tank) MoveUp() {
	v := t.GetPos()
	vn := pixel.V(v.X, v.Y+5)
	t.SetPos(&vn)
	// t.pos.X = t.pos.X - 5
	t.SetOrientation(interfaces.OrientationTop)
}

// MoveDown передвигает танк вправо
func (t *Tank) MoveDown() {
	v := t.GetPos()
	vn := pixel.V(v.X, v.Y-5)
	t.SetPos(&vn)
	t.SetOrientation(interfaces.OrientationBottom)
}
