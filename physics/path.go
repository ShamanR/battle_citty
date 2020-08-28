package physics

import (
	"github.com/beefsack/go-astar"
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
)

func (p *Physics) PathTo(from, to pixel.Vec, sceneMap consts.LevelMap) []*pixel.Vec {
	fromTile := newTile(from, sceneMap, p.tileSize, p.scale)
	toTile := newTile(to, sceneMap, p.tileSize, p.scale)

	path, _, found := astar.Path(fromTile, toTile)
	if !found {
		return nil
	}

	res := make([]*pixel.Vec, len(path))
	for _, pp := range path {
		t := pp.(*tile)
		v := pixel.V(float64(t.x * p.tileSize * p.scale), float64(t.y * p.tileSize * p.scale))
		res = append(res, &v)
	}
	return res
}
