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
	life        uint8 // Количество жизней у объекта

	// sprite
	spriteList  *interfaces.SceneObjectAnimateList
	spriteIndex int64 // 0, 1, 2...

	// child
	children []interfaces.SceneObject

	gameObject interface{}
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
		children: []interfaces.SceneObject{},
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
	drawListener, ok := o.GetGameObject().(interfaces.DrawListener)
	if ok {
		drawListener.OnDraw()
	}
	s := o.GetSprite()
	if s == nil || !o.IsVisible() {
		return
	}
	if o.objectType == consts.ObjectTypeProjectile {
		//fmt.Println("here")
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

func (o *Object) AddChild(obj interfaces.SceneObject) {
	o.children = append(o.children, obj)
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
	o.scene.RemoveObject(o.id)
}

func (o *Object) GetGameObject() interface{} {
	return o.gameObject
}

func (o *Object) SetGameObject(gObj interface{}) {
	o.gameObject = gObj
}

func (o *Object) SetLife(life uint8) {
	o.life = life
}

func (o *Object) GetLife() uint8 {
	return o.life
}

// Метод вызывается при столкновении с другим объектом сцены
func (o *Object) OnCollide(with interfaces.SceneObject) {
	if firstDestructable, ok := o.gameObject.(interfaces.Damageable); ok {
		firstDestructable.OnDamage(with)
	}

	if secondDestroyable, ok := with.GetGameObject().(interfaces.Damageable); ok {
		secondDestroyable.OnDamage(o)
	}
}
