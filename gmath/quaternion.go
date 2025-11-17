package gmath

import "math"

type Quaternion struct {
	W, X, Y, Z float32
}

func (q Quaternion) Normalize() Quaternion {
	mag := float32(math.Sqrt(float64(q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z)))
	return Quaternion{
		W: q.W / mag,
		X: q.X / mag,
		Y: q.Y / mag,
		Z: q.Z / mag,
	}
}

func (q Quaternion) Invert() Quaternion {
	c := q.Normalize()
	return Quaternion{
		W: c.W,
		X: -c.X,
		Y: -c.Y,
		Z: -c.Z,
	}
}

func (q Quaternion) Multiply(o Quaternion) Quaternion {
	return Quaternion{
		W: q.W*o.W - q.X*o.X - q.Y*o.Y - q.Z*o.Z,
		X: q.W*o.X + q.X*o.W + q.Y*o.Z - q.Z*o.Y,
		Y: q.W*o.Y - q.X*o.Z + q.Y*o.W + q.Z*o.X,
		Z: q.W*o.Z + q.X*o.Y - q.Y*o.X + q.Z*o.W,
	}
}

func QuaternionFromAxisAngle(axis Vector3, angle float32) Quaternion {
	half := angle * 0.5
	s := float32(math.Sin(float64(half)))

	axis = axis.Normalize()

	return Quaternion{
		W: float32(math.Cos(float64(half))),
		X: axis.X * s,
		Y: axis.Y * s,
		Z: axis.Z * s,
	}.Normalize()
}
