package camera

import (
	"math"
	render_math "opal/render_math"
)

// The Camera interface is used for rendering to
// generate a ViewMatrix and ProjectionMatrix
// for rendering purposes. It is implemented
// by BasicCamera.
type Camera interface {
	ViewMatrix() render_math.Matrix4
	ProjectionMatrix() render_math.Matrix4
}

// ------

type BasicCamera struct {
	Position render_math.Vector3

	FOV, Aspect, Near, Far float32
	Rotation               render_math.Quaternion
}

func NewBasicCamera(w, h int) *BasicCamera {
	return &BasicCamera{
		Position: render_math.Vector3{},
		Rotation: render_math.Quaternion{W: 1},
		FOV:      70 * (math.Pi / 180),
		Aspect:   float32(w) / float32(h),
		Near:     0.1,
		Far:      100.0,
	}
}

func (b *BasicCamera) ViewMatrix() render_math.Matrix4 {
	// Use inverse rotation for view matrix
	r := b.Rotation.Normalize().Invert()

	// Precalculate repeated terms
	xx := r.X * r.X
	yy := r.Y * r.Y
	zz := r.Z * r.Z
	xy := r.X * r.Y
	xz := r.X * r.Z
	yz := r.Y * r.Z
	wx := r.W * r.X
	wy := r.W * r.Y
	wz := r.W * r.Z

	// Build rotation part
	rot := render_math.Matrix4{M: [16]float32{
		1 - 2*(yy+zz), 2 * (xy + wz), 2 * (xz - wy), 0,
		2 * (xy - wz), 1 - 2*(xx+zz), 2 * (yz + wx), 0,
		2 * (xz + wy), 2 * (yz - wx), 1 - 2*(xx+yy), 0,
		0, 0, 0, 1,
	}}

	// Translation: apply inverse translation
	invPos := render_math.Vector3{X: -b.Position.X, Y: -b.Position.Y, Z: -b.Position.Z}

	// Multiply rotation * translation
	// Equivalent to: rot * Translate(-position)
	rot.M[12] = rot.M[0]*invPos.X + rot.M[4]*invPos.Y + rot.M[8]*invPos.Z
	rot.M[13] = rot.M[1]*invPos.X + rot.M[5]*invPos.Y + rot.M[9]*invPos.Z
	rot.M[14] = rot.M[2]*invPos.X + rot.M[6]*invPos.Y + rot.M[10]*invPos.Z

	return rot
}

func (b *BasicCamera) ProjectionMatrix() render_math.Matrix4 {
	return render_math.PerspectiveMatrix4(b.FOV, b.Aspect, b.Near, b.Far)
}

func (b *BasicCamera) MoveForward(m float32) {
	forward := render_math.Vector3{Z: -1}.Rotate(b.Rotation)
	b.Position = b.Position.Add(forward.Scale(m))
}

func (b *BasicCamera) MoveRight(m float32) {
	right := render_math.Vector3{X: 1}.Rotate(b.Rotation)
	b.Position = b.Position.Add(right.Scale(m))
}

func (b *BasicCamera) MoveUp(m float32) {
	up := render_math.Vector3{Y: 1}.Rotate(b.Rotation)
	b.Position = b.Position.Add(up.Scale(m))
}

func (b *BasicCamera) Yaw(rad float32) {
	// Always yaw around world up
	rot := render_math.QuaternionFromAxisAngle(render_math.Vector3{Y: 1}, rad)
	b.Rotation = rot.Multiply(b.Rotation).Normalize()
}

func (b *BasicCamera) Pitch(rad float32) {
	// Pitch around camera's local right
	right := render_math.Vector3{X: 1}.Rotate(b.Rotation)
	rot := render_math.QuaternionFromAxisAngle(right, rad)
	b.Rotation = rot.Multiply(b.Rotation).Normalize()
}

func (b *BasicCamera) Roll(rad float32) {
	forward := render_math.Vector3{Z: -1}.Rotate(b.Rotation)
	rot := render_math.QuaternionFromAxisAngle(forward, rad)
	b.Rotation = rot.Multiply(b.Rotation).Normalize()
}
