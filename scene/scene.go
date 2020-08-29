package scene

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
	object "github.com/shamanr/battle_citty/scene/objects"
)

type Scene struct {
	objects    []interfaces.SceneObject
	objCounter int64
	level      consts.LevelMap
	spawner interfaces.Spawner
}

func NewScene(spawner interfaces.Spawner) *Scene {
	return &Scene{
		objects:    []interfaces.SceneObject{},
		objCounter: 0,
		spawner: spawner,
	}
}

func (s *Scene) GetObjects() []interfaces.SceneObject {
	return s.objects
}

func (s *Scene) GetObjectByID(id int64) interfaces.SceneObject {
	for _, obj := range s.objects {
		if obj.GetID() == id {
			return obj
		}
	}
	return nil
}

func (s *Scene) RemoveObject(id int64) {
	for i, obj := range s.objects {
		if obj.GetID() == id {
			s.objects = append(s.objects[:i], s.objects[i+1:]...)
		}
	}
}

func (s *Scene) GetSceneMap() interfaces.SceneMap {
	return s.objects
}

func (s *Scene) GetLevelMap() consts.LevelMap {
	return s.level
}

func (s *Scene) SetLevelMap(level consts.LevelMap) {
	s.level = level
}

func (s *Scene) Draw(target pixel.Target) {
	for _, obj := range s.objects {
		obj.Draw(target)
	}
}

func (s *Scene) MakeEmptyObj(objType consts.ObjectType) interfaces.SceneObject {
	id := s.objCounter
	s.objCounter++
	pos := pixel.V(-100, -100) // за пределами экрана
	obj := object.NewObject(id, s, objType, &pos, nil)
	obj.SetVisible(false)
	s.objects = append(s.objects, obj)
	return obj
}

func (s *Scene) SetSceneObjects(objects []interfaces.SceneObject) {
	s.objects = objects
	s.objCounter = int64(len(objects))
}

func (s *Scene) GetSpawner() interfaces.Spawner {
	return s.spawner
}

