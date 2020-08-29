package actors

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/scene/objects/tank"
	"math/rand"
	"time"
)

type AI struct {
	*tank.Tank
	rand               *rand.Rand
	prevPosition       *pixel.Vec
	currentOrientation consts.Orientation
	wait               time.Duration
}

func NewAI() *AI {
	return &AI{
		currentOrientation: consts.OrientationBottom,
		prevPosition:       &pixel.ZV,
	}
}

func (ai *AI) SetTank(tank *tank.Tank) {
	ai.Tank = tank
	ai.rand = rand.New(rand.NewSource(time.Now().Unix()))
}

func (ai *AI) Tick(dt time.Duration) {
	ai.RandomShoot()
	if ai.wait < time.Millisecond*300 {
		ai.wait += dt
		ai.Move(ai.currentOrientation)
		return
	}
	ai.wait = 0
	if ai.Tank.GetPos().Eq(*ai.prevPosition) { // Танк затрял
		ai.RandomDirection()
		return
	}
	ai.Move(ai.currentOrientation)
	ai.prevPosition = ai.Tank.GetPos()
}

func (ai *AI) Move(orientation consts.Orientation) {
	ai.currentOrientation = orientation
	switch orientation {
	case consts.OrientationTop:
		ai.Tank.MoveUp()
	case consts.OrientationLeft:
		ai.Tank.MoveLeft()
	case consts.OrientationBottom:
		ai.Tank.MoveDown()
	case consts.OrientationRight:
		ai.Tank.MoveRight()
	}
}

func (ai *AI) RandomDirection() {
	switch ai.rand.Intn(4) {
	case consts.OrientationTop:
		ai.Move(consts.OrientationTop)
	case consts.OrientationLeft:
		ai.Move(consts.OrientationLeft)
	case consts.OrientationBottom:
		ai.Move(consts.OrientationBottom)
	case consts.OrientationRight:
		ai.Move(consts.OrientationRight)
	}
}

func (ai *AI) RandomShoot() {
	if ai.rand.Intn(100) > 90 {
		ai.Tank.Shoot()
	}
}
