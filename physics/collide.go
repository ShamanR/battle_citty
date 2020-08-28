package physics

import (
	i "github.com/shamanr/battle_citty/interfaces"
)

func (p *Physics) initCollisionsMap() {
	p.collisionRules = map[i.ObjectType]map[i.ObjectType]bool{
		i.ObjectTypePlayerTank1: {
			i.ObjectTypePlayerTank1: true,
			i.ObjectTypeBrickWall:   true,
			i.ObjectTypeIronWall:    true,
			i.ObjectTypeWater:       true,
			i.ObjectTypeProjectile:  true,
		},
		i.ObjectTypeProjectile: {
			i.ObjectTypePlayerTank1:  true,
			i.ObjectTypeBrickWall:    true,
			i.ObjectTypeIronWall:     true,
			i.ObjectTypeHeadquarters: true,
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
