package physics

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/internal/scene/objects"
	"testing"
)

func Test_isMoving(t *testing.T) {
	tests := []struct {
		name string
		vec  pixel.Vec
		want bool
	}{
		{
			name: "Двигается!",
			vec:  pixel.V(10000, -3000),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := object.Object{}
			obj.SetSpeed(&tt.vec)
			if got := isMoving(&obj); got != tt.want {
				t.Errorf("isMoving() = %v, want %v", got, tt.want)
			}
		})
	}
}
