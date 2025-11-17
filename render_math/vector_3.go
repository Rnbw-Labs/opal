package render_math

import "math"

type Vector3 struct {
	X, Y, Z float32
}

func (v Vector3) Add(o Vector3) Vector3 {
	return Vector3{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v Vector3) Sub(o Vector3) Vector3 {
	return Vector3{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

func (v Vector3) Negative() Vector3 {
	return Vector3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vector3) Scale(s float32) Vector3 {
	return Vector3{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

func (v Vector3) Dot(o Vector3) float32 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector3) Cross(o Vector3) Vector3 {
	return Vector3{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X,
	}
}

func (v Vector3) Length() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vector3) Normalize() Vector3 {
	l := v.Length()
	if l == 0 {
		return v
	}
	return v.Scale(1 / l)
}

func Translate(v Vector3) Matrix4 {
	m := IdentityMatrix4()
	m.M[12] = v.X
	m.M[13] = v.Y
	m.M[14] = v.Z
	return m
}

func LookAt(eye, target, up Vector3) Matrix4 {
	f := target.Sub(eye).Normalize()
	s := f.Cross(up.Normalize()).Normalize()
	u := s.Cross(f)
	return Matrix4{
		M: [16]float32{
			s.X, u.X, -f.X, 0,
			s.Y, u.Y, -f.Y, 0,
			s.Z, u.Z, -f.Z, 0,
			-s.Dot(eye), -u.Dot(eye), f.Dot(eye), 1,
		},
	}
}

func (v Vector3) Rotate(q Quaternion) Vector3 {
	// Convert v into a pure quaternion
	vx, vy, vz := v.X, v.Y, v.Z

	// Precalculate repeated operations
	x2 := q.X * 2
	y2 := q.Y * 2
	z2 := q.Z * 2
	xx2 := q.X * x2
	yy2 := q.Y * y2
	zz2 := q.Z * z2
	xy2 := q.X * y2
	xz2 := q.X * z2
	yz2 := q.Y * z2
	wx2 := q.W * x2
	wy2 := q.W * y2
	wz2 := q.W * z2

	return Vector3{
		X: (1-(yy2+zz2))*vx + (xy2-wz2)*vy + (xz2+wy2)*vz,
		Y: (xy2+wz2)*vx + (1-(xx2+zz2))*vy + (yz2-wx2)*vz,
		Z: (xz2-wy2)*vx + (yz2+wx2)*vy + (1-(xx2+yy2))*vz,
	}
}
