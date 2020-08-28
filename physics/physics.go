package physics

import (
	"github.com/shamanr/battle_citty/consts"
	"time"
)

type Physics struct {
	frameDuration  time.Duration
	collisionRules map[consts.ObjectType]map[consts.ObjectType]bool
}

func New(frameDuration time.Duration) *Physics {
	p := &Physics{
		frameDuration: frameDuration,
	}
	p.initCollisionsMap()
	return p
}
