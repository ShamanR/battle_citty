package scene

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/interfaces"
	object "github.com/shamanr/battle_citty/scene/objects"
)

type Scene struct {
	objects    []interfaces.SceneObject
	objCounter int64
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

func (s *Scene) AddObject(object interfaces.SceneObject) {
	s.objects = append(s.objects, object)
}

func (s *Scene) GetSceneMap() interfaces.SceneMap {
	return s.objects
}

func (s *Scene) GetLevelMap() interfaces.LevelMap {
	panic("implement me")
}

func (s *Scene) Draw(target pixel.Target) {
	for _, obj := range s.objects {
		obj.Draw(target)
	}
}

func (s *Scene) MakeEmptyObj() interfaces.SceneObject {
	id := s.objCounter
	s.objCounter++
	pos := pixel.V(-100, -100) // за пределами экрана
	obj := object.NewObject(id, interfaces.ObjectTypeEmpty, &pos, nil)
	obj.SetVisible(false)
	return obj
}
