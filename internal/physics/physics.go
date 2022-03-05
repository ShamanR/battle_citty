package physics

import (
	"github.com/shamanr/battle_citty/internal/consts"
	"time"
)

type Physics struct {
	frameDuration  time.Duration
	tileSize       int
	scale          int
	collisionRules map[consts.ObjectType]map[consts.ObjectType]bool
}

func New(frameDuration time.Duration, tileSize, scale int) *Physics {
	p := &Physics{
		frameDuration: frameDuration,
		tileSize:      tileSize,
		scale:         scale,
	}
	p.initCollisionsMap()
	return p
}
