package physics_test

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/internal/consts"
	"github.com/shamanr/battle_citty/internal/physics"
	"testing"
)

func TestPhysics_PathTo(t *testing.T) {
	tests := []struct {
		name            string
		tileSize, scale int
		from, to        pixel.Vec
		level           func() consts.LevelMap
		expectations    func(t *testing.T, p []*pixel.Vec)
	}{
		{
			name:     "test",
			tileSize: 1,
			scale:    1,
			from:     pixel.V(0, 0),
			to:       pixel.V(1, 1),
			level: func() consts.LevelMap {
				level := make(consts.LevelMap, 3)
				for i := range level {
					level[i] = make([]consts.ObjectType, 3)
				}
				return level
			},
			expectations: func(t *testing.T, p []*pixel.Vec) {
				if len(p) != 3 {
					t.Errorf("path is too long: %v", p)
				}

			},
		},
		{
			name:     "test2",
			tileSize: 1,
			scale:    1,
			from:     pixel.V(0, 0),
			to:       pixel.V(0, 2),
			level: func() consts.LevelMap {
				level := make(consts.LevelMap, 3)
				for i := range level {
					level[i] = make([]consts.ObjectType, 3)
				}
				level[1][0] = consts.ObjectTypeBrickWall
				level[1][1] = consts.ObjectTypeBrickWall
				return level
			},
			expectations: func(t *testing.T, p []*pixel.Vec) {
				if len(p) != 7 {
					t.Errorf("path is too long: %v", p)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := physics.New(1, test.tileSize, test.scale)
			path := p.PathTo(test.from, test.to, test.level())
			test.expectations(t, path)
		})
	}
}
