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

func (l *Loader) GetTankAnimation(spriteRow int) []*pixel.Sprite {
	return l.loadAnimation(0, 256-spriteRow*16, 16, 8)
}

func (l *Loader) loadAnimation(topX int, topY int, spriteSize int, frames int) []*pixel.Sprite {
	result := make([]*pixel.Sprite, 0, frames)
	for i := 0; i < frames; i++ {
		sprite := pixel.NewSprite(
			l.texture,
			pixel.Rect{
				pixel.Vec{float64(topX + spriteSize*i), float64(topY)},
				pixel.Vec{float64(topX + spriteSize*(i+1)), float64(topY - spriteSize)},
			},
		)
		result = append(result, sprite)
	}
	return result
}

func (l *Loader) MakeUserTank() *tank.Tank {
	obj := objects.Object{}
	obj.SetPosition(pixel.ZV)
	obj.Show()
	obj.SetAnimation(l.GetTankAnimation(3))
	obj.SetScale(3)

	animation := l.GetTankAnimation(3)
	tankAnimation := map[int][]*pixel.Sprite{
		tank.OrientationTop:   {animation[0], animation[1]},
		tank.OrientationLeft:  {animation[2], animation[3]},
		tank.OrientationDown:  {animation[4], animation[5]},
		tank.OrientationRight: {animation[6], animation[7]},
	}
	return &tank.Tank{Object: obj, TankAnimation: tankAnimation}
}
