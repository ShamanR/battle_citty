package physics

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
)

func (p *Physics) PathTo(from, to pixel.Vec, sceneMap consts.LevelMap) {
	fromTile := newTile(from, sceneMap)
	toTile := newTile(to, sceneMap)

	path, distance, found := astar.Path(fromTile, toTile)
}
