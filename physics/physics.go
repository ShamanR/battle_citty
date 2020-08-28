package physics

import (
	i "github.com/shamanr/battle_citty/interfaces"
	"time"
)

type Physics struct {
	frameDuration time.Duration
	collisionRules map[i.ObjectType]map[i.ObjectType]bool
}

func New(frameDuration time.Duration) *Physics {
	p := &Physics{}
	p.initCollisionsMap()
	return p
}
