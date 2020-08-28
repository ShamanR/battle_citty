package object

import (
	"github.com/faiface/pixel"
	"github.com/shamanr/battle_citty/consts"
	"github.com/shamanr/battle_citty/interfaces"
)

// Object базовая структура
type Object struct {
	id         int64
	scene      interfaces.Scene
	objectType consts.ObjectType

	// props
	visible bool
	scale   pixel.Vec

	// bounds and pos
	pos         pixel.Vec
	orientation consts.Orientation
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
func NewObject(ID int64, scene interfaces.Scene, objectType consts.ObjectType, pos *pixel.Vec, spriteList *interfaces.SceneObjectAnimateList) *Object {
	obj := Object{
		id:         ID,
		scene:      scene,
		objectType: objectType,
		bounds: pixel.Rect{
			Min: pixel.Vec{0, 0},
			Max: pixel.Vec{16 * consts.ScaleSprites, 16 * consts.ScaleSprites},
		},
	}
	obj.SetPos(pos)
	obj.SetSpriteList(spriteList)
	return &obj
}

func (o *Object) GetID() int64 {
	return o.id
}

func (o *Object) GetScene() interfaces.Scene {
	return o.scene
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

// GetSprite возвращает активный спрайт
func (o *Object) GetSprite() *pixel.Sprite {
	if o.spriteList == nil || len(*o.spriteList) == 0 {
		return nil
	}
	return (*o.spriteList)[o.orientation][o.spriteIndex]
}

func (o *Object) NextSprite() {
	o.spriteIndex++
	o.spriteIndex %= int64(len((*o.spriteList)[o.orientation]))
}

// SetSpriteList обновляет spriteList объекта
func (o *Object) SetSpriteList(list *interfaces.SceneObjectAnimateList) {
	o.spriteList = list
}

// Draw выполняет отрисовку объекта в target
func (o *Object) Draw(target pixel.Target) {
	s := o.GetSprite()
	if s == nil || !o.IsVisible() {
		return
	}
	m := pixel.IM.ScaledXY(pixel.ZV, o.scale).Moved(*o.GetPos())
	s.Draw(target, m)
	for _, nestedObj := range o.GetObjects() {
		nestedObj.Draw(target)
	}
}

// SetScale устанавливает коэф. масштабирования объекта
func (o *Object) SetScale(scale pixel.Vec) {
	o.scale = scale
	o.bounds = pixel.R(0, 0, consts.MapTileSize*scale.X, consts.MapTileSize*scale.Y)
}

// GetScale возвращает коэф. масштабирования объекта
func (o *Object) GetScale() pixel.Vec {
	return o.scale
}

// Bounds возвращает границы объекта
func (o *Object) GetSize() *pixel.Rect {
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
func (o *Object) GetObjectType() consts.ObjectType {
	return o.objectType
}

// SetOrientation выставляет ориентацию объекта
func (o *Object) SetOrientation(orient consts.Orientation) {
	o.orientation = orient
}

func (o *Object) GetOrientation() consts.Orientation {
	return o.orientation
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

func (o *Object) onCollide(obj interfaces.SceneObject) {
}
