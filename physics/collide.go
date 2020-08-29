package physics

import (
	"github.com/shamanr/battle_citty/consts"
	i "github.com/shamanr/battle_citty/interfaces"
)

func (p *Physics) initCollisionsMap() {
	tank := map[consts.ObjectType]bool{
		consts.ObjectTypePlayerTank1: true,
		consts.ObjectTypePlayerTank2: true,
		consts.ObjectTypePlayerTank3: true,
		consts.ObjectTypePlayerTank4: true,
		consts.ObjectTypeEnemyTank1:  true,
		consts.ObjectTypeEnemyTank2:  true,
		consts.ObjectTypeEnemyTank3:  true,
		consts.ObjectTypeEnemyTank4:  true,
		consts.ObjectTypeBrickWall:   true,
		consts.ObjectTypeIronWall:    true,
		consts.ObjectTypeWater:       true,
		consts.ObjectTypeProjectile:  true,
		consts.ObjectTypeBrickWallDamagedTop: true,
		consts.ObjectTypeBrickWallDamagedLeft: true,
		consts.ObjectTypeBrickWallDamagedRight: true,
		consts.ObjectTypeBrickWallDamagedDown: true,
		consts.ObjectTypeIronWallDamagedTop: true,
		consts.ObjectTypeIronWallDamagedLeft: true,
		consts.ObjectTypeIronWallDamagedRight: true,
		consts.ObjectTypeIronWallDamagedDown: true,
		consts.ObjectTypeHiddenWall: true,
	}

	p.collisionRules = map[consts.ObjectType]map[consts.ObjectType]bool{
		consts.ObjectTypePlayerTank1: tank,
		consts.ObjectTypePlayerTank2: tank,
		consts.ObjectTypePlayerTank3: tank,
		consts.ObjectTypePlayerTank4: tank,
		consts.ObjectTypeEnemyTank1:  tank,
		consts.ObjectTypeEnemyTank2:  tank,
		consts.ObjectTypeEnemyTank3:  tank,
		consts.ObjectTypeEnemyTank4:  tank,
		consts.ObjectTypeProjectile: {
			consts.ObjectTypePlayerTank1:  true,
			consts.ObjectTypePlayerTank2:  true,
			consts.ObjectTypePlayerTank3:  true,
			consts.ObjectTypePlayerTank4:  true,
			consts.ObjectTypeEnemyTank1:   true,
			consts.ObjectTypeEnemyTank2:   true,
			consts.ObjectTypeEnemyTank3:   true,
			consts.ObjectTypeEnemyTank4:   true,
			consts.ObjectTypeBrickWall:    true,
			consts.ObjectTypeIronWall:     true,
			consts.ObjectTypeHeadquarters: true,
			consts.ObjectTypeBrickWallDamagedTop: true,
			consts.ObjectTypeBrickWallDamagedLeft: true,
			consts.ObjectTypeBrickWallDamagedRight: true,
			consts.ObjectTypeBrickWallDamagedDown: true,
			consts.ObjectTypeIronWallDamagedTop: true,
			consts.ObjectTypeIronWallDamagedLeft: true,
			consts.ObjectTypeIronWallDamagedRight: true,
			consts.ObjectTypeIronWallDamagedDown: true,
			consts.ObjectTypeHiddenWall: true,
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
