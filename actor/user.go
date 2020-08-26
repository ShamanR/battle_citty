package actor

import (
	"github.com/faiface/pixel/pixelgl"
	"time"
)

type Tank interface {
	MoveLeft()
	MoveRight()
	MoveTop()
	MoveDown()
	Shoot()
}

type User struct {
	Tank Tank
}

type keyboardInterface interface {
	Pressed(button pixelgl.Button) bool
}

func (u *User) AttachToKeyboard(keyboard keyboardInterface) {
	for {
		if keyboard.Pressed(pixelgl.KeyLeft) {
			u.Tank.MoveLeft()
		}
		if keyboard.Pressed(pixelgl.KeyRight) {
			u.Tank.MoveRight()
		}
		if keyboard.Pressed(pixelgl.KeyUp) {
			u.Tank.MoveTop()
		}
		if keyboard.Pressed(pixelgl.KeyUp) {
			u.Tank.MoveDown()
		}
		<-time.After(time.Millisecond * 30)
	}
}
