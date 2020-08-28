package tank

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/interfaces"
)

// NewTank возвращает объект танка
func NewTank(obj interfaces.SceneObject) *Tank {
	return &Tank{obj}
}

// Tank структруа танка
type Tank struct {
	interfaces.SceneObject
}

// MoveLeft передвигает танк влево
func (t *Tank) MoveLeft() {
	t.move()
	t.SetOrientation(interfaces.OrientationLeft)
}

// MoveRight передвигает танк вправо
func (t *Tank) MoveRight() {
	t.move()
	t.SetOrientation(interfaces.OrientationRight)
}

// MoveUp передвигает танк влево
func (t *Tank) MoveUp() {
	t.move()
	t.SetOrientation(interfaces.OrientationTop)
}

// MoveDown передвигает танк вправо
func (t *Tank) MoveDown() {
	t.move()
	t.SetOrientation(interfaces.OrientationBottom)
}

func (t *Tank) move() {
	v := t.GetPos()
	s := t.GetSpeed()
	vn := pixel.V(v.X+s.X, v.Y+s.Y)
	t.SetPos(&vn)
}

// Stop остановка танка
func (t *Tank) Stop() {
	s := pixel.V(0, 0)
	t.SetSpeed(&s)
}

// Shoot стрельба
func (t *Tank) Shoot() {
	// TODO:
}
