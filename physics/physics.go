package physics

import i "github.com/shamanr/battle_citty/interfaces"

type Physics struct {
	collisionRules map[i.ObjectType]map[i.ObjectType]bool
}

func New() *Physics {
	return &Physics{}
}
