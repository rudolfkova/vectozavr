package object

import (
	"github.com/rudolfkova/vectozavr"
)

type Object struct {
	position          vectozavr.Vec3
	transformMatrix   vectozavr.Matrix
	angle             vectozavr.Vec3
	angleLeftUpLookAt vectozavr.Vec3
	left              vectozavr.Vec3
	up                vectozavr.Vec3
	lookAt            vectozavr.Vec3
}

func (o *Object) Transform(t vectozavr.Matrix) {
	o.transformMatrix = o.transformMatrix.MatMul(t)
}

func (o *Object) Left() {
	o.left = o.transformMatrix.X()
}

func (o *Object) Up() {
	o.up = o.transformMatrix.Y()
}

func (o *Object) LookAt() {
	o.lookAt = o.transformMatrix.Z()
}

func (o *Object) TransformRelativePoint(point vectozavr.Vec3, transform vectozavr.Matrix) {
	// translate object in new coordinate system (connected with point)
	o.transformMatrix = vectozavr.Translation(o.position.Sub(point)).MatMul(o.transformMatrix)
	// transform object in the new coordinate system
	o.transformMatrix = transform.MatMul(o.transformMatrix)
	// translate object back in self connected coordinate system
	o.position = o.transformMatrix.W().Add(point)
	o.transformMatrix = vectozavr.Translation(o.transformMatrix.W()).MatMul(o.transformMatrix)

}

func (o *Object) Translate(v vectozavr.Vec3) {
	o.position = o.position.Add(v)
}

func (o *Object) Scale(s vectozavr.Vec3) {
	o.Transform(vectozavr.Scale(s))
}

func (o *Object) Rotate(a vectozavr.Vec3) {
	o.angle = o.angle.Add(a)
	o.Transform(vectozavr.Rotation(a))
}

func (o *Object) VRotate(v vectozavr.Vec3, a float64) {
	o.Transform(vectozavr.RotationV(v, a))
}

func (o *Object) RotateRelativePoint(s vectozavr.Vec3, r vectozavr.Vec3) {
	o.angle = o.angle.Add(r)
	o.TransformRelativePoint(s, vectozavr.Rotation(r))
}

func (o *Object) RotateLeft(rl float64) {
	o.angleLeftUpLookAt.X += rl
	o.VRotate(o.left, rl)
}

func (o *Object) RotateUp(rl float64) {
	o.angleLeftUpLookAt.Y += rl
	o.VRotate(o.up, rl)
}

func (o *Object) RotateLookAt(rl float64) {
	o.angleLeftUpLookAt.Z += rl
	o.VRotate(o.lookAt, rl)
}

func (o *Object) TranslateToPoint(point vectozavr.Vec3) {
	o.Translate(point.Sub(o.position))
}
