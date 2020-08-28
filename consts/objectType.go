package consts

// ObjectType тип объекта
type ObjectType uint8

// LevelMap ...
type LevelMap [][]ObjectType

const (
	ObjectTypeEmpty ObjectType = iota
	ObjectTypeBrickWall
	ObjectTypeIronWall
	ObjectTypeWater
	ObjectTypeForest
	ObjectTypeIce
	ObjectTypeHeadquarters
	ObjectTypePlayerSpawn
	ObjectTypeAISpawn
	ObjectTypeBonus
	ObjectTypePlayerTank1
	ObjectTypePlayerTank2
	ObjectTypePlayerTank3
	ObjectTypePlayerTank4
	ObjectTypeEnemyTank1
	ObjectTypeEnemyTank2
	ObjectTypeEnemyTank3
	ObjectTypeEnemyTank4
	ObjectTypeProjectile
)
