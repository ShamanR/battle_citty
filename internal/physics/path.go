package physics

import (
	"github.com/beefsack/go-astar"
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/internal/consts"
)

func (p *Physics) PathTo(from, to pixel.Vec, sceneMap consts.LevelMap) []*pixel.Vec {
	tc := newTileContainer(sceneMap)

	fromTile := tc.getTile(
		int(from.X/float64(p.tileSize*p.scale)),
		int(from.Y/float64(p.tileSize*p.scale)),
	)
	toTile := tc.getTile(
		int(to.X/float64(p.tileSize*p.scale)),
		int(to.Y/float64(p.tileSize*p.scale)),
	)

	path, _, found := astar.Path(toTile, fromTile)
	if !found {
		return nil
	}

	res := make([]*pixel.Vec, 0, len(path))
	for _, pp := range path {
		t := pp.(*tile)
		v := pixel.V(float64(t.x*p.tileSize*p.scale), float64(t.y*p.tileSize*p.scale))
		res = append(res, &v)
	}
	return res
}
