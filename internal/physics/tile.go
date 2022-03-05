package physics

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"github.com/shamanr/battle_citty/internal/consts"
)

type tile struct {
	x, y      int
	Container *tileContainer
}

type tileContainer struct {
	level consts.LevelMap
	tiles [][]*tile
}

func newTileContainer(level consts.LevelMap) *tileContainer {
	tiles := make([][]*tile, len(level))
	for i := range tiles {
		tiles[i] = make([]*tile, len(level[i]))
	}
	return &tileContainer{
		level: level,
		tiles: tiles,
	}
}

func (c *tileContainer) getTile(x, y int) *tile {
	if x < 0 || y < 0 || y >= len(c.level) || x >= len(c.level[0]) {
		return nil
	}

	t := c.tiles[y][x]
	if t != nil {
		return t
	}

	t = &tile{
		x:         x,
		y:         y,
		Container: c,
	}
	c.tiles[y][x] = t
	return t
}

func (t *tile) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		nx := t.x + offset[0]
		ny := t.y + offset[1]

		nt := t.Container.getTile(nx, ny)
		if nt == nil {
			continue
		}

		if !isFreeType((t.Container.level)[ny][nx]) {
			continue
		}

		neighbors = append(neighbors, nt)
	}
	return neighbors
}

func (t *tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t *tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*tile)
	absX := toT.x - t.x
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - t.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func (t tile) String() string {
	return fmt.Sprintf("%d.%d", t.x, t.y)
}

func isFreeType(objectType consts.ObjectType) bool {
	return objectType != consts.ObjectTypeBrickWall && objectType != consts.ObjectTypeIronWall && objectType != consts.ObjectTypeWater
}
