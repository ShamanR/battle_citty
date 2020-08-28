package resource_manager

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
)

var sides = [4]consts.Orientation{
	consts.OrientationTop,
	consts.OrientationLeft,
	consts.OrientationBottom,
	consts.OrientationRight,
}

type animationPosition struct {
	spriteSheetSize pixel.Vec
	spriteSize      int
	positionX       int
	positionY       int
	frames          int
	movable         bool
}

func newAnimationPosition(spriteSheetSize pixel.Vec, spriteSize, startX, startY, frames int, movable bool) *animationPosition {
	return &animationPosition{
		spriteSheetSize: spriteSheetSize,
		spriteSize:      spriteSize,
		positionX:       startX,
		positionY:       startY,
		frames:          frames,
		movable:         movable,
	}
}
