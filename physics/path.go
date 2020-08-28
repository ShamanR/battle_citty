package physics

import (
	"github.com/beefsack/go-astar"
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
)

func (p *Physics) PathTo(from, to pixel.Vec, sceneMap consts.LevelMap) {
	fromTile := newTile(from, sceneMap)
	toTile := newTile(to, sceneMap)

	//path, distance, found := astar.Path(fromTile, toTile)
	_, _, _ = astar.Path(fromTile, toTile)
}
