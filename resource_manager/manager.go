package resource_manager

import (
	"github.com/faiface/pixel"
	"github.com/pkg/errors"
	"github.com/shamanr/battle_citty/interfaces"
	"image"
	"os"
)

const (
	spriteSheetSizeY = 256
	spriteSheetSizeX = 400
	defaultSpriteSize = 16

	SimpleTankOrangeUp     interfaces.SpriteType = "SimpleTankOrangeUp"
	SimpleTankOrangeUpMove            = "SimpleTankOrangeUpMove"
)

var spriteSheetSize = pixel.V(spriteSheetSizeX, spriteSheetSizeY)

var spriteMap = map[interfaces.SpriteType]*spritePosition{
	:     newSpritePosition(spriteSheetSize, defaultSpriteSize, 0, 0),
	SimpleTankOrangeUpMove: newSpritePosition(spriteSheetSize, defaultSpriteSize, 1, 0),
}

//var animationsMap = map[interfaces.ObjectType]*interfaces.SceneObjectAnimateList{
//	interfaces.SimpleOrangeTank:
//}
//
//type animation struct {
//
//}

type spritePosition struct {
	spriteSheetSize pixel.Vec
	spriteSize int
	positionX int
	positionY int
}

func newSpritePosition(spriteSheetSize pixel.Vec, size int, posX int, posY int) *spritePosition {
	return &spritePosition{
		spriteSheetSize: spriteSheetSize,
		spriteSize: size,
		positionX:  posX,
		positionY:  posY,
	}
}

func (s *spritePosition) Bounds() pixel.Rect {
	spriteStartY := s.spriteSheetSize.Y - float64(s.positionY * s.spriteSize)
	spriteStartX := float64(s.positionX * s.spriteSize)

	return pixel.R(spriteStartX, spriteStartY, spriteStartX + float64(s.spriteSize), spriteStartY - float64(s.spriteSize))
}

type resourceManager struct {
	spriteSheet pixel.Picture
	cache map[interfaces.SpriteType]*pixel.Sprite
}

func NewResourceManager(spritePath string) *resourceManager {
	spriteSheet, err := loadPicture(spritePath)
	if err != nil {
		panic(errors.Wrap(err, "Unable to load spriteSheet"))
	}

	return &resourceManager{
		spriteSheet: spriteSheet,
		cache: make(map[interfaces.SpriteType]*pixel.Sprite),
	}
}

func (s *resourceManager) GetSprite(name interfaces.SpriteType) *pixel.Sprite {
	if sprite, ok := s.cache[name]; ok {
		return sprite
	}

	spriteElement, ok := spriteMap[name]
	if !ok {
		panic(errors.New("Unable to find sprite by name"))
	}

	sprite := pixel.NewSprite(s.spriteSheet, spriteElement.Bounds())
	s.cache[name] = sprite

	return sprite
}

func (s *resourceManager) GetSpriteMap(name interfaces.ObjectType) *interfaces.SceneObjectAnimateList {

}

func (s *resourceManager) LoadMap() *interfaces.SceneMap {
	return &interfaces.SceneMap{}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
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
