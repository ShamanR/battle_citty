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
	speedVec := t.getOrientationVec()
	speedVec.Scaled(float64(t.tankSpeed))
	t.SetSpeed(speedVec)
	t.NextSprite()
}

func (t *Tank) getOrientationVec() *pixel.Vec {
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
	return &vec
}

// Stop остановка танка
func (t *Tank) Stop() {
	s := pixel.V(0, 0)
	t.SetSpeed(&s)
}

// Shoot стрельба
func (t *Tank) Shoot() {
	bulletObj := t.GetScene().MakeEmptyObj(consts.ObjectTypeProjectile)
	vec := t.GetPos().Add((*t.getOrientationVec()).Scaled(48))
	bulletObj.SetPos(&vec)
	bulletObj.SetOrientation(t.GetOrientation())
	bulletObj.SetVisible(true)
	bulletObj.SetScale(t.GetScale().Scaled(0.5))
	spriteList := interfaces.SceneObjectAnimateList{}
	sprite := t.GetSprite()
	spriteList[t.GetOrientation()] = []*pixel.Sprite{sprite}
	bulletObj.SetSpriteList(&spriteList)
	speed := t.GetSpeed().Scaled(2)
	bulletObj.SetSpeed(&speed)
}
