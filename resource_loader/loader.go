package resource_loader

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/scene/objects"
	"github.com/shamanr/battle_citty/scene/objects/tank"
	"image"
	_ "image/png"
	"os"
)

type Loader struct {
	ResourcePath string
	texture      pixel.Picture
}

func (l *Loader) loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(l.ResourcePath + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (l *Loader) Init() {
	textures, err := l.loadPicture("textures.png")
	if err != nil {
		panic("error in loader " + err.Error())
	}
	l.texture = textures
}

func (l *Loader) GetUserTankAnimation() []*pixel.Sprite {
	size := float64(16)
	sprite1 := pixel.NewSprite(l.texture, pixel.Rect{pixel.Vec{0, 256}, pixel.Vec{size, 256 - size}})
	sprite2 := pixel.NewSprite(l.texture, pixel.Rect{pixel.Vec{size, 256}, pixel.Vec{size * 2, 256 - size}})
	return []*pixel.Sprite{
		sprite1,
		sprite2,
	}
}

func (l *Loader) MakeUserTank() *tank.UserTank {
	obj := objects.Object{}
	obj.SetPosition(pixel.ZV)
	obj.Show()
	obj.SetAnimation(l.GetUserTankAnimation())
	obj.SetScale(3)
	return &tank.UserTank{obj}
}
