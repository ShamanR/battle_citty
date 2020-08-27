package objects

import "github.com/faiface/pixel"

type Object struct {
	speed          pixel.Vec
	position       pixel.Vec
	angle          float64
	scale          float64
	visible        bool
	animation      []*pixel.Sprite
	animationFrame int
	nestedObjects  []*Object
}

func (o *Object) SetSpeed(vec pixel.Vec) {
	o.speed = vec
}

func (o *Object) GetSpeed() pixel.Vec {
	return o.speed
}

func (o *Object) GetPosition() pixel.Vec {
	return o.position
}

func (o *Object) SetPosition(pos pixel.Vec) {
	o.position = pos
}

func (o *Object) Rotate(angle float64) {
	o.angle += angle
}

func (o *Object) GetAngle() float64 {
	return o.angle
}

func (o *Object) SetAngle(angle float64) {
	o.angle = angle
}

func (o *Object) IsVisible() bool {
	return o.visible
}

func (o *Object) Hide() {
	o.visible = false
}

func (o *Object) Show() {
	o.visible = true
}
func (o *Object) SetScale(scale float64) {
	o.scale = scale
}

func (o *Object) GetCurrentSprite() *pixel.Sprite {
	return o.animation[o.animationFrame]
}

func (o *Object) SetAnimation(animation []*pixel.Sprite) {
	o.animation = animation
	o.animationFrame = 0
}

func (o *Object) NextAnimationFrame() {
	o.animationFrame++
	if o.animationFrame > len(o.animation)-1 {
		o.animationFrame = 0
	}
}

func (o *Object) Draw(target pixel.Target) {
	if o.IsVisible() {
		sprite := o.GetCurrentSprite()
		mat := pixel.IM.Rotated(pixel.ZV, o.angle).Scaled(pixel.ZV, o.scale).Moved(o.GetPosition())
		sprite.Draw(target, mat)
	}
	for _, nestedObject := range o.nestedObjects {
		nestedObject.Draw(target)
	}
}
