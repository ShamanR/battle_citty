package physics

import (
	"github.com/beefsack/go-astar"
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
)

func (p *Physics) PathTo(from, to pixel.Vec, sceneMap consts.LevelMap) []*pixel.Vec {
	fromTile := newTile(from, sceneMap)
	toTile := newTile(to, sceneMap)

	path, _, found := astar.Path(fromTile, toTile)
	if !found {
		return nil
	}

	res := make([]*pixel.Vec, len(path))
	for _, p := range path {
		t := p.(*tile)
		v := pixel.V(float64(t.x * tileSize), float64(t.y * tileSize))
		res = append(res, &v)
	}
	return res
}
