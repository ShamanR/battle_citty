package physics

import (
	"github.com/beefsack/go-astar"
	"github.com/faiface/pixel"
	i "github.com/shamanr/battle_citty/interfaces"
)

func (p *Physics) PathTo(from, to pixel.Vec, sceneMap i.LevelMap) {
	fromTile := newTile(from, sceneMap)
	toTile := newTile(to, sceneMap)

	path, distance, found := astar.Path(fromTile, toTile)
}
