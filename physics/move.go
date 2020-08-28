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

		movement := calcMovement(currentObj, dt, p.frameDuration)
		newBounds := bounds(currentObj).Moved(movement)

		for _, anotherObj := range sceneMap {
			if anotherObj == currentObj {
				continue
			}

			if !p.areColliable(currentObj, anotherObj) {
				continue
			}

			if newBounds.Intersects(bounds(anotherObj)) {
				continue loop
			}
		}

		newPos := currentObj.GetPos().Add(movement)
		currentObj.SetPos(&newPos)
	}
}

func calcMovement(currentObj interfaces.SceneObject, dt, fd time.Duration) pixel.Vec {
	frames := dt / fd
	movement := currentObj.GetSpeed().Scaled(float64(frames))
	return movement
}

func isMoving(object interfaces.SceneObject) bool {
	return !object.GetSpeed().Eq(pixel.ZV)
}

func bounds(obj interfaces.SceneObject) pixel.Rect {
	b := obj.GetSize().Moved(*obj.GetPos())
	return b.Resized(b.Center(), pixel.V(b.W()*0.96, b.H()*0.96))
}
