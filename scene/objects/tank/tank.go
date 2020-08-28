package tank

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
)

// NewTank возвращает объект танка
func NewTank(obj interfaces.SceneObject, speed int) *Tank {
	return &Tank{
		SceneObject: obj,
		tankSpeed:   speed,
	}
}

// Tank структруа танка
type Tank struct {
	interfaces.SceneObject
	tankSpeed int
}

// MoveLeft передвигает танк влево
func (t *Tank) MoveLeft() {
	t.SetOrientation(consts.OrientationLeft)
	t.move()
}

// MoveRight передвигает танк вправо
func (t *Tank) MoveRight() {
	t.SetOrientation(consts.OrientationRight)
	t.move()
}

// MoveUp передвигает танк влево
func (t *Tank) MoveUp() {
	t.move()
	t.SetOrientation(consts.OrientationTop)
}

// MoveDown передвигает танк вправо
func (t *Tank) MoveDown() {
	t.SetOrientation(consts.OrientationBottom)
	t.move()
}

func (t *Tank) move() {
	var vec pixel.Vec
	switch t.GetOrientation() {
	case consts.OrientationTop:
		vec.Y = 1 * float64(t.tankSpeed)
	case consts.OrientationLeft:
		vec.X = -1 * float64(t.tankSpeed)
	case consts.OrientationBottom:
		vec.Y = -1 * float64(t.tankSpeed)
	case consts.OrientationRight:
		vec.X = 1 * float64(t.tankSpeed)
	}
	t.SetSpeed(&vec)
	t.NextSprite()
}

// Stop остановка танка
func (t *Tank) Stop() {
	s := pixel.V(0, 0)
	t.SetSpeed(&s)
}

// Shoot стрельба
func (t *Tank) Shoot() {

}
