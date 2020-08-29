package object

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
)

// NewMovableObject возвращает объект танка
func NewMovableObject(obj interfaces.SceneObject, speed int) *MovableObject {
	return &MovableObject{
		SceneObject: obj,
		speed:       speed,
	}
}

// MovableObject структруа танка
type MovableObject struct {
	interfaces.SceneObject
	speed int
}

// MoveLeft передвигает танк влево
func (t *MovableObject) MoveLeft() {
	t.SetOrientation(consts.OrientationLeft)
	t.move()
}

// MoveRight передвигает танк вправо
func (t *MovableObject) MoveRight() {
	t.SetOrientation(consts.OrientationRight)
	t.move()
}

// MoveUp передвигает танк влево
func (t *MovableObject) MoveUp() {
	t.move()
	t.SetOrientation(consts.OrientationTop)
}

// MoveDown передвигает танк вправо
func (t *MovableObject) MoveDown() {
	t.SetOrientation(consts.OrientationBottom)
	t.move()
}

func (t *MovableObject) move() {
	speedVec := t.getOrientationVec()
	speedVec.Scaled(float64(t.speed))
	t.SetSpeed(speedVec)
	t.NextSprite()
}

func (t *MovableObject) getOrientationVec() *pixel.Vec {
	var vec pixel.Vec
	switch t.GetOrientation() {
	case consts.OrientationTop:
		vec.Y = 1 * float64(t.speed)
	case consts.OrientationLeft:
		vec.X = -1 * float64(t.speed)
	case consts.OrientationBottom:
		vec.Y = -1 * float64(t.speed)
	case consts.OrientationRight:
		vec.X = 1 * float64(t.speed)
	}
	return &vec
}

// Stop остановка танка
func (t *MovableObject) Stop() {
	s := pixel.V(0, 0)
	t.SetSpeed(&s)
}
