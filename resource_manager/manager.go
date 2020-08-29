package resource_manager

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/pkg/errors"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
)

type SpriteType string

const (
	spriteSheetSizeY  = 256
	spriteSheetSizeX  = 400
	defaultSpriteSize = 16

	SimpleTankOrangeUp     SpriteType = "SimpleTankOrangeUp"
	SimpleTankOrangeUpMove            = "SimpleTankOrangeUpMove"
)

var spriteSheetSize = pixel.V(spriteSheetSizeX, spriteSheetSizeY)

var spriteMap = map[SpriteType]*spritePosition{
	SimpleTankOrangeUp:     newSpritePosition(spriteSheetSize, defaultSpriteSize, 0, 0),
	SimpleTankOrangeUpMove: newSpritePosition(spriteSheetSize, defaultSpriteSize, 1, 0),
}

var animationsMap = map[consts.ObjectType]*animationPosition{
	consts.ObjectTypePlayerTank1:           newAnimationPosition(spriteSheetSize, defaultSpriteSize, 0, 0, 2, true),
	consts.ObjectTypeBrickWall:             newAnimationPosition(spriteSheetSize, defaultSpriteSize, 16, 0, 1, false),
	consts.ObjectTypeBrickWallDamagedLeft:  newAnimationPosition(spriteSheetSize, defaultSpriteSize, 17, 0, 1, false),
	consts.ObjectTypeBrickWallDamagedTop:   newAnimationPosition(spriteSheetSize, defaultSpriteSize, 18, 0, 1, false),
	consts.ObjectTypeBrickWallDamagedRight: newAnimationPosition(spriteSheetSize, defaultSpriteSize, 19, 0, 1, false),
	consts.ObjectTypeBrickWallDamagedDown: newAnimationPosition(spriteSheetSize, defaultSpriteSize, 20, 0, 1, false),
	consts.ObjectTypeIronWall:   newAnimationPosition(spriteSheetSize, defaultSpriteSize, 16, 1, 1, false),
	consts.ObjectTypeIronWallDamagedLeft: newAnimationPosition(spriteSheetSize, defaultSpriteSize, 17, 1, 1, false),
	consts.ObjectTypeIronWallDamagedTop: newAnimationPosition(spriteSheetSize, defaultSpriteSize, 18, 1, 1, false),
	consts.ObjectTypeIronWallDamagedRight: newAnimationPosition(spriteSheetSize, defaultSpriteSize, 19, 1, 1, false),
	consts.ObjectTypeIronWallDamagedDown: newAnimationPosition(spriteSheetSize, defaultSpriteSize, 20, 1, 1, false),
	consts.ObjectTypeHeadquarters:   newAnimationPosition(spriteSheetSize, defaultSpriteSize, 19, 2, 1, false),
	consts.ObjectTypePlayerSpawn: newAnimationPosition(spriteSheetSize, defaultSpriteSize, 16, 6, 4, false),
	consts.ObjectTypeProjectile: newAnimationPosition(spriteSheetSize, 4, 20, 6, 1, true),
	consts.ObjectTypeExplosion:             newAnimationPosition(spriteSheetSize, defaultSpriteSize, 16, 8, 3, false),
}

type spritePosition struct {
	spriteSheetSize pixel.Vec
	spriteSize      int
	positionX       int
	positionY       int
	startX			int
	startY			int
}

func newSpritePosition(spriteSheetSize pixel.Vec, size int, posX int, posY int) *spritePosition {
	return &spritePosition{
		spriteSheetSize: spriteSheetSize,
		spriteSize:      size,
		positionX:       posX,
		positionY:       posY,
		startX:			 defaultSpriteSize*posX,
		startY:			 spriteSheetSizeY - defaultSpriteSize*posY,
	}
}

func (s *spritePosition) Bounds() pixel.Rect {
	//mult := 0
	//if s.spriteSize < defaultSpriteSize {
	//	mult = defaultSpriteSize-s.spriteSize
	//}
	return pixel.R(float64(s.startX), float64(s.startY), float64(s.startX+s.spriteSize), float64(s.startY-s.spriteSize))
}

type resourceManager struct {
	spriteSheet pixel.Picture
	cache       map[SpriteType]*pixel.Sprite
}

func NewResourceManager(spritePath string) *resourceManager {
	spriteSheet, err := loadPicture(spritePath)
	if err != nil {
		panic(errors.Wrap(err, "Unable to load spriteSheet"))
	}

	return &resourceManager{
		spriteSheet: spriteSheet,
		cache:       make(map[SpriteType]*pixel.Sprite),
	}
}

func (s *resourceManager) GetSprite(name SpriteType) *pixel.Sprite {
	spriteElement, ok := spriteMap[name]
	if !ok {
		panic(errors.New("Unable to find sprite by name"))
	}

	return s.loadSprite(name, spriteElement)
}

func (rm *resourceManager) getSceneObjectAnimateList(name consts.ObjectType, position *animationPosition) *interfaces.SceneObjectAnimateList {
	animationsList := &interfaces.SceneObjectAnimateList{}
	// 325 104
	// 332 103
	// 341 104
	// 348 104
	//if name == consts.ObjectTypeProjectile {
	//	(*animationsList)[consts.OrientationTop] = []*pixel.Sprite{
	//		rm.loadSprite(),
	//	}
	//}


	framesCounter := 0
	currentSide := 0
	framesBuff := make([]*pixel.Sprite, 0, position.frames)
	for i := 0; i < position.frames*len(sides); i++ {
		spritePos := newSpritePosition(position.spriteSheetSize, position.spriteSize, position.positionX+i, position.positionY)

		if name == consts.ObjectTypeProjectile {
			spritePos = newSpritePosition(position.spriteSheetSize, position.spriteSize, 1, 1)
			y := spriteSheetSizeY - 104
			if i == consts.OrientationTop {
				spritePos.startX = 321
				spritePos.startY = y
			} else if i == consts.OrientationRight {
				spritePos.startX = 328
				spritePos.startY = y
			} else if i == consts.OrientationBottom {
				spritePos.startX = 337
				spritePos.startY = y
			} else if i == consts.OrientationRight {
				spritePos.startX = 344
				spritePos.startY = y
			}
		}

		sprite := rm.loadSprite(SpriteType(fmt.Sprintf("%d%d", spritePos.positionY, spritePos.positionX)), spritePos)
		framesBuff = append(framesBuff, sprite)

		framesCounter++
		if framesCounter == position.frames {
			(*animationsList)[sides[currentSide]] = framesBuff
			if !position.movable {
				return animationsList
			}
			framesCounter = 0
			currentSide++
			framesBuff = make([]*pixel.Sprite, 0, position.frames)
		}
	}

	return animationsList
}

func (s *resourceManager) GetSpriteMap(name consts.ObjectType) *interfaces.SceneObjectAnimateList {
	animationPos, ok := animationsMap[name]
	if !ok {
		panic(errors.New(fmt.Sprintf("Unable to load animations by name %f", name)))
	}

	return s.getSceneObjectAnimateList(name, animationPos)
}

func (s *resourceManager) loadSprite(id SpriteType, spriteElement *spritePosition) *pixel.Sprite {
	if sprite, ok := s.cache[id]; ok {
		return sprite
	}

	sprite := pixel.NewSprite(s.spriteSheet, spriteElement.Bounds())
	s.cache[id] = sprite

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
