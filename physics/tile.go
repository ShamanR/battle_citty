package physics

import (
	"github.com/beefsack/go-astar"
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
)

const tileSize = 16

type tile struct {
	x, y int
	level consts.LevelMap
}

func newTile(pos pixel.Vec, level consts.LevelMap) tile {
	return tile{
		x: int(pos.X / tileSize),
		y: int(pos.X / tileSize),
		level: level,
	}
}

func (t tile) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		nx := t.x + offset[0]
		ny := t.y + offset[1]

		if !isInsideLevel(nx, ny, t.level) {
			continue
		}

		if !isFreeType(t.level[ny][nx]) {
			continue
		}

		neighbors = append(neighbors, newTile(pixel.V(float64(nx), float64(ny)), t.level))
	}
	return neighbors
}

func (t tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*tile)
	absX := toT.x - t.y
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - t.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func isFreeType(objectType consts.ObjectType) bool {
	return objectType != consts.ObjectTypeBrickWall && objectType != consts.ObjectTypeIronWall && objectType != consts.ObjectTypeWater
}

func isInsideLevel(x, y int, level consts.LevelMap) bool {
	return x < 0 || y < 0 || x >= len(level[0]) || y >= len(level)
}

