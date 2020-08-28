package actors

import (
	"github.com/faiface/pixel/pixelgl"
)

type Tank interface {
	MoveLeft()
	MoveRight()
	MoveUp()
	MoveDown()
	Stop()
	Shoot()
}

type User struct {
	Tank Tank
}

type keyboardInterface interface {
	Pressed(button pixelgl.Button) bool
}

func (u *User) SetTank(tank Tank) {
	u.Tank = tank
}

func (u *User) AttachToKeyboard(keyboard keyboardInterface) {
	if keyboard.Pressed(pixelgl.KeyLeft) {
		u.Tank.MoveLeft()
		return
	}
	if keyboard.Pressed(pixelgl.KeyRight) {
		u.Tank.MoveRight()
		return
	}
	if keyboard.Pressed(pixelgl.KeyUp) {
		u.Tank.MoveUp()
		return
	}
	if keyboard.Pressed(pixelgl.KeyDown) {
		u.Tank.MoveDown()
		return
	}
	u.Tank.Stop()
}
