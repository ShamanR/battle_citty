package physics

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/interfaces"
	"time"
)

func (p *Physics) MoveObjects(sceneMap interfaces.SceneMap, dt time.Duration) {
loop:
	for _, currentObj := range sceneMap {
		if !isMoving(currentObj) {
			continue
		}

		movement := currentObj.GetSpeed().Scaled(float64(dt))
		newBounds := currentObj.Bounds().Moved(movement)

		for _, anotherObj := range sceneMap {
			if anotherObj == currentObj {
				continue
			}

			if !p.canCollide(currentObj, anotherObj) {
				continue
			}

			if newBounds.Intersects(*anotherObj.Bounds()) {
				continue loop
			}
		}

		newPos := currentObj.GetPos().Add(movement)
		currentObj.SetPos(&newPos)
	}
}

func isMoving(object interfaces.SceneObject) bool {
	return object.GetSpeed().Eq(pixel.Vec{X: 0, Y: 0})
}
