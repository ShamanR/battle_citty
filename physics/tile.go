package physics

import (
	"github.com/beefsack/go-astar"
	"github.com/faiface/pixel"
	i "github.com/shamanr/battle_citty/interfaces"
)

type tile struct {
	x, y int
	level i.LevelMap
}

func newTile(pos pixel.Vec, level i.LevelMap) tile {
	return tile{
		x: int(pos.X / i.TileSize),
		y: int(pos.X / i.TileSize),
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

func isFreeType(objectType i.ObjectType) bool {
	return objectType != i.ObjectTypeBrickWall && objectType != i.ObjectTypeIronWall && objectType != i.ObjectTypeWater
}

func isInsideLevel(x, y int, level i.LevelMap) bool {
	return x < 0 || y < 0 || x >= len(level[0]) || y >= len(level)
}

