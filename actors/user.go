package actors

import (
	"github.com/faiface/pixel/pixelgl"
	"time"
)

type Tank interface {
	MoveLeft()
	MoveRight()
	MoveTop()
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

func (u *User) AttachToKeyboard(keyboard keyboardInterface) {
	for {
		<-time.After(time.Millisecond * 30)
		if keyboard.Pressed(pixelgl.KeyLeft) {
			u.Tank.MoveLeft()
			continue
		}
		if keyboard.Pressed(pixelgl.KeyRight) {
			u.Tank.MoveRight()
			continue
		}
		if keyboard.Pressed(pixelgl.KeyUp) {
			u.Tank.MoveTop()
			continue
		}
		if keyboard.Pressed(pixelgl.KeyDown) {
			u.Tank.MoveDown()
			continue
		}
		u.Tank.Stop()
	}
}
