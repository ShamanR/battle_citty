package physics

import (
	i "github.com/shamanr/battle_citty/interfaces"
)

func (p *Physics) initCollisionsMap() {
	tank := map[i.ObjectType]bool{
		i.ObjectTypePlayerTank1: true,
		i.ObjectTypePlayerTank2: true,
		i.ObjectTypePlayerTank3: true,
		i.ObjectTypePlayerTank4: true,
		i.ObjectTypeEnemyTank1:  true,
		i.ObjectTypeEnemyTank2:  true,
		i.ObjectTypeEnemyTank3:  true,
		i.ObjectTypeEnemyTank4:  true,
		i.ObjectTypeBrickWall:   true,
		i.ObjectTypeIronWall:    true,
		i.ObjectTypeWater:       true,
		i.ObjectTypeProjectile:  true,
	}

	p.collisionRules = map[i.ObjectType]map[i.ObjectType]bool{
		i.ObjectTypePlayerTank1: tank,
		i.ObjectTypePlayerTank2: tank,
		i.ObjectTypePlayerTank3: tank,
		i.ObjectTypePlayerTank4: tank,
		i.ObjectTypeEnemyTank1:  tank,
		i.ObjectTypeEnemyTank2:  tank,
		i.ObjectTypeEnemyTank3:  tank,
		i.ObjectTypeEnemyTank4:  tank,
		i.ObjectTypeProjectile: {
			i.ObjectTypePlayerTank1:  true,
			i.ObjectTypePlayerTank2:  true,
			i.ObjectTypePlayerTank3:  true,
			i.ObjectTypePlayerTank4:  true,
			i.ObjectTypeEnemyTank1:   true,
			i.ObjectTypeEnemyTank2:   true,
			i.ObjectTypeEnemyTank3:   true,
			i.ObjectTypeEnemyTank4:   true,
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
