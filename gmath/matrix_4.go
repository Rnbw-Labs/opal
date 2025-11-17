package gmath

import "math"

type Matrix4 struct {
	M [16]float32
}

func IdentityMatrix4() Matrix4 {
	return Matrix4{[16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}}
}

func (a Matrix4) Multiply(b Matrix4) Matrix4 {
	var r Matrix4
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			sum := float32(0)
			for k := 0; k < 4; k++ {
				sum += a.M[row+4*k] * b.M[k+4*col]
			}
			r.M[row+4*col] = sum
		}
	}
	return r
}

func PerspectiveMatrix4(fov, aspect, near, far float32) Matrix4 {
	f := 1.0 / float32(math.Tan(float64(fov)/2))
	nf := 1 / (near - far)

	return Matrix4{[16]float32{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (far + near) * nf, -1,
		0, 0, (2 * far * near) * nf, 0,
	}}
}
