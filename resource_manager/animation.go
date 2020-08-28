package resource_manager

import "github.com/shamanr/battle_citty/interfaces"

func NewSceneObjectAnimateList(startX, startY, frames int) *interfaces.SceneObjectAnimateList {
	for i := 0; i < frames * 4; i++ {
		sprite := newSpritePosition(spriteSheetSize, defaultSpriteSize, startX+i, startY+i)
	}

	return &interfaces.SceneObjectAnimateList{
		LeftSprite:   nil,
		RightSprite:  nil,
		TopSprite:    nil,
		BottomSprite: nil,
	}
}


