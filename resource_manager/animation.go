package resource_manager

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/interfaces"
)

var sides = [4]interfaces.Orientation{
	interfaces.OrientationTop,
	interfaces.OrientationLeft,
	interfaces.OrientationBottom,
	interfaces.OrientationRight,
}

type animationPosition struct {
	spriteSheetSize pixel.Vec
	spriteSize int
	positionX int
	positionY int
	frames int
	movable bool
}

func newAnimationPosition(spriteSheetSize pixel.Vec, spriteSize, startX, startY, frames int, movable bool) *animationPosition {
	return &animationPosition{
		spriteSheetSize: spriteSheetSize,
		spriteSize: spriteSize,
		positionX:  startX,
		positionY:  startY,
		frames: frames,
		movable: movable,
	}
}


