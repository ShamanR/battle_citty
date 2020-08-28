package physics

import (
	i "github.com/shamanr/battle_citty/interfaces"
)

func (p *Physics) initCollisionsMap() {
	p.collisionRules = map[i.ObjectType]map[i.ObjectType]bool{
		i.SimpleOrangeTank: {
			i.SimpleOrangeTank: true,
			i.BrickWall:        true,
			i.IronWall:         true,
			i.Water:            true,
			i.Projectile:       true,
		},
		i.Projectile: {
			i.SimpleOrangeTank: true,
			i.BrickWall:        true,
			i.IronWall:         true,
			i.Headquarters:     true,
		},
	}
}

func (p *Physics) areColliable(obj, obj2 i.SceneObject) bool {
	v, ok := p.collisionRules[obj.GetObjectType()]
	if ok {
		return v[obj2.GetObjectType()]
	}

	v, ok = p.collisionRules[obj2.GetObjectType()]
	if !ok {
		return v[obj.GetObjectType()]
	}

	return false
}
