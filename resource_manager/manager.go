package resource_manager

import (
	"github.com/faiface/pixel"
	"github.com/pkg/errors"
	"image"
	"os"
)

type ObjectType uint8
//type SceneMap []SceneObject
type LevelMap [][]ObjectType
type SpriteType string

const (
	BrickWall ObjectType = iota
	IronWall
	Water
	Forest
	Ice
	Headquarters
	PlayerSpawn
	AISpawn
	Bonus
)

const (
	spriteSheetSizeY = 400
	spriteSheetSizeX = 256
	defaultSpriteSize = 16

	SimpleTankOrangeUp     SpriteType = "SimpleTankOrangeUp"
	SimpleTankOrangeUpMove            = "SimpleTankOrangeUpMove"
)

var spriteSheetSize = pixel.V(spriteSheetSizeX, spriteSheetSizeY)

var spriteMap = map[SpriteType]*SpritePosition{
	SimpleTankOrangeUp:     newSpritePosition(spriteSheetSize, defaultSpriteSize, 1, 1),
	SimpleTankOrangeUpMove: newSpritePosition(spriteSheetSize, defaultSpriteSize, 2, 1),
}

type SpritePosition struct {
	spriteSheetSize pixel.Vec
	spriteSize float64
	positionX float64
	positionY float64
}

func newSpritePosition(spriteSheetSize pixel.Vec, size float64, posX float64, posY float64) *SpritePosition {
	return &SpritePosition{
		spriteSheetSize: spriteSheetSize,
		spriteSize: size,
		positionX:  posX,
		positionY:  posY,
	}
}

func (s *SpritePosition) Bounds() pixel.Rect {
	spriteStartY := s.spriteSheetSize.Y - s.positionX * s.spriteSize
	spriteStartX := s.spriteSheetSize.X + s.positionY * s.spriteSize

	return pixel.R(spriteStartX, spriteStartY, spriteStartX + s.spriteSize, spriteStartY + s.spriteSize)
}

type resourceManager struct {
	spriteSheet pixel.Picture
	cache map[SpriteType]*pixel.Sprite
}

func NewResourceManager(spritePath string) *resourceManager {
	spriteSheet, err := loadPicture(spritePath)
	if err != nil {
		panic(errors.Wrap(err, "Unable to load spriteSheet"))
	}

	return &resourceManager{
		spriteSheet: spriteSheet,
		cache: make(map[SpriteType]*pixel.Sprite),
	}
}

func (s *resourceManager) GetSprite(name SpriteType) *pixel.Sprite {
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
