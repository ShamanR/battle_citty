package object

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/interfaces"
)

// Object базовая структура
type Object struct {
	id         int64
	objectType interfaces.ObjectType

	// props
	visible bool
	scale   float64

	// bounds and pos
	pos         pixel.Vec
	orientation interfaces.Orientation
	angle       float64 // возможно не нужно пока или рассчитывать автоматически по orientation
	speed       pixel.Vec
	bounds      pixel.Rect

	// sprite
	spriteList  *interfaces.SceneObjectAnimateList
	spriteIndex int64 // 0, 1, 2...

	// child
	children []interfaces.SceneObject
}

// NewObject конструктор:
// - objectType -- тип объекта
// - pos -- позиция объекта на карте
// - spriteList -- структура спрайтов для анимации
func NewObject(ID int64, objectType interfaces.ObjectType, pos *pixel.Vec, spriteList *interfaces.SceneObjectAnimateList) *Object {
	obj := Object{
		id:         ID,
		objectType: objectType,
	}
	obj.SetPos(pos)
	obj.SetSpriteList(spriteList)
	return &obj
}

func (o *Object) GetID() int64 {
	return o.id
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
	o.spriteIndex++
	o.spriteIndex %= int64(len((*o.spriteList)[o.orientation]))
	return (*o.spriteList)[o.orientation][o.spriteIndex]
}

// SetSpriteList обновляет spriteList объекта
func (o *Object) SetSpriteList(list *interfaces.SceneObjectAnimateList) {
	o.spriteList = list
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
func (o *Object) GetObjects() []interfaces.SceneObject {
	// return o.children
	return o.children
}

// GetObjectType возвращает тип объекта
func (o *Object) GetObjectType() interfaces.ObjectType {
	return o.objectType
}

// SetOrientation выставляет ориентацию объекта
func (o *Object) SetOrientation(orient interfaces.Orientation) {
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
