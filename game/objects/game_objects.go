package objects

import (
	"errors"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
	"github.com/shamanr/battle_citty/scene/objects"
	"github.com/shamanr/battle_citty/scene/objects/projectile"
	"github.com/shamanr/battle_citty/scene/objects/tank"
)

var gameObjectsDefaultConfigs = map[consts.ObjectType]*GameObjectConfig{
	consts.ObjectTypePlayerTank1: newGameObjectConfig(consts.ObjectTypePlayerTank1, 1, 5),
	consts.ObjectTypeProjectile:  newGameObjectConfig(consts.ObjectTypeProjectile, 1, 1),
}

type GameObjectsManager struct {
	rm interfaces.ResourceManager
	scene interfaces.Scene
}

func NewGameObjectsManager(rm interfaces.ResourceManager, scene interfaces.Scene) *GameObjectsManager {
	return &GameObjectsManager{
		rm:    rm,
		scene: scene,
	}
}

func (manager *GameObjectsManager) MakeTank(typ consts.ObjectType) *tank.Tank {
	config, ok := gameObjectsDefaultConfigs[typ]
	if !ok {
		panic(errors.New("Unable to find tank config"))
	}

	obj := manager.scene.MakeEmptyObj(typ)
	obj.SetSpriteList(manager.rm.GetSpriteMap(typ))
	obj.SetVisible(true)
	sceneTank := object.NewMovableObject(obj, config.Speed)
	bullet := manager.MakeProjectile(consts.ObjectTypeProjectile)

	return tank.NewTank(sceneTank, bullet)
}

func (manager *GameObjectsManager) MakeProjectile(typ consts.ObjectType) *projectile.Projectile {
	config, ok := gameObjectsDefaultConfigs[typ]
	if !ok {
		panic(errors.New("Unable to find projectile config"))
	}

	obj := manager.scene.MakeEmptyObj(typ)
	obj.SetSpriteList(manager.rm.GetSpriteMap(typ))
	obj.SetVisible(true)
	sceneProjectile := object.NewMovableObject(obj, config.Speed)

	return projectile.NewProjectile(sceneProjectile)
}

type GameObjectConfig struct {
	Typ consts.ObjectType
	Life uint8
	Speed int
}

func newGameObjectConfig(typ consts.ObjectType, Life uint8, Speed int) *GameObjectConfig {
	return &GameObjectConfig{
		Typ: typ,
		Life:  Life,
		Speed: Speed,
	}
}

