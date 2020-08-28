package object

import (
	"github.com/faiface/pixel"
)

// ObjectType тип объекта
type ObjectType uint8
type Orientation uint8

const (
	OrientationTop    = 0
	OrientationRight  = 1
	OrientationBottom = 2
	OrientationLeft   = 3
)

// type SceneMap []SceneObject
// type LevelMap [][]ObjectType

// SceneObjectAnimateList структуры анимации (для танков?)
type SceneObjectAnimateList struct {
	LeftSprite   []*pixel.Sprite
	RightSprite  []*pixel.Sprite
	TopSprite    []*pixel.Sprite
	BottomSprite []*pixel.Sprite
}

// SceneObject интерфейс
type SceneObject interface {
	GetPos() *pixel.Vec
	SetPos(vect *pixel.Vec)
	GetSpeed() *pixel.Vec
	SetSpeed(vect *pixel.Vec)
	getSprite() *pixel.Sprite
	SetSpriteList(list *SceneObjectAnimateList)
	getSriteListLen() int64
	Draw(target pixel.Target)
	SetScale(scale float64)
	GetScale() float64
	Bounds() *pixel.Rect
	IsVisible() bool
	SetVisible(visible bool)
	GetObjects() []SceneObject
	GetObjectType() ObjectType
	SetOrientation(orient Orientation)
	Delete()
}

// Object базовая структура
type Object struct {
	objectType ObjectType

	// props
	visible bool
	scale   float64

	// bounds and pos
	pos         pixel.Vec
	orientation Orientation
	angle       float64 // возможно не нужно пока или рассчитывать автоматически по orientation
	speed       pixel.Vec
	bounds      pixel.Rect

	// sprite
	spriteList  *SceneObjectAnimateList
	spriteIndex int64 // 0, 1, 2...

	// child
	children []SceneObject
}

// NewObject конструктор:
// - objectType -- тип объекта
// - pos -- позиция объекта на карте
// - spriteList -- структура спрайтов для анимации
func NewObject(objectType ObjectType, pos *pixel.Vec, spriteList *SceneObjectAnimateList) *Object {
	obj := Object{
		objectType: objectType,
	}
	obj.SetPos(pos)
	obj.SetSpriteList(spriteList)
	return &obj
}

// GetPos возвращает позицию объекта
func (o *Object) GetPos() *pixel.Vec {
	return &o.pos
}

// SetPos устанавливает позицию объекта
func (o *Object) SetPos(vect *pixel.Vec) {
	o.pos = pixel.V(vect.X, vect.Y)
}

// GetSpeed возвращает вектор скорости объекта
func (o *Object) GetSpeed() *pixel.Vec {
	return &o.speed
}

// SetSpeed устанавливает вектор скорости объекта
func (o *Object) SetSpeed(vect *pixel.Vec) {
	o.speed = pixel.V(vect.X, vect.Y)
}

// getSprite возвращает активный спрайт
func (o *Object) getSprite() *pixel.Sprite {
	o.spriteIndex += 1
	o.spriteIndex %= o.getSriteListLen()
	if o.orientation == OrientationTop {
		return o.spriteList.TopSprite[o.spriteIndex]
	}
	if o.orientation == OrientationRight {
		return o.spriteList.RightSprite[o.spriteIndex]
	}
	if o.orientation == OrientationBottom {
		return o.spriteList.BottomSprite[o.spriteIndex]
	}
	if o.orientation == OrientationLeft {
		return o.spriteList.LeftSprite[o.spriteIndex]
	}
	// TODO: возможно тут понадобится другая логика
	return nil
}

// SetSpriteList обновляет spriteList объекта
func (o *Object) SetSpriteList(list *SceneObjectAnimateList) {
	o.spriteList = list
}

func (o *Object) getSriteListLen() int64 {
	l := 0
	if o.orientation == OrientationTop {
		l = len(o.spriteList.TopSprite)
	}
	if o.orientation == OrientationRight {
		l = len(o.spriteList.RightSprite)
	}
	if o.orientation == OrientationBottom {
		l = len(o.spriteList.BottomSprite)
	}
	if o.orientation == OrientationLeft {
		l = len(o.spriteList.LeftSprite)
	}
	return int64(l)
}

// Draw выполняет отрисовку объекта в target
func (o *Object) Draw(target pixel.Target) {
	s := o.getSprite()
	m := pixel.IM.Scaled(pixel.ZV, o.scale).Moved(*o.GetPos())
	s.Draw(target, m)
}

// SetScale устанавливает коэф. масштабирования объекта
func (o *Object) SetScale(scale float64) {
	o.scale = scale
}

// GetScale возвращает коэф. масштабирования объекта
func (o *Object) GetScale() float64 {
	return o.scale
}

// Bounds возвращает границы объекта
func (o *Object) Bounds() *pixel.Rect {
	return &o.bounds
}

// IsVisible возвращает текущую видимость объекта
func (o *Object) IsVisible() bool {
	return o.visible
}

// SetVisible устанавливает видимость объекта
func (o *Object) SetVisible(visible bool) {
	o.visible = visible
}

// GetObjects возвращает жочерние объекты
func (o *Object) GetObjects() []SceneObject {
	// return o.children
	return o.children
}

// GetObjectType возвращает тип объекта
func (o *Object) GetObjectType() ObjectType {
	return o.objectType
}

// SetOrientation выставляет ориентацию объекта
func (o *Object) SetOrientation(orient Orientation) {
	o.orientation = orient
}

// Delete TODO: вроде как так должно работать
func (o *Object) Delete() {
	if o.children != nil {
		for _, ch := range o.children {
			ch.Delete()
		}
		o.children = nil

	}
}
