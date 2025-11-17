package gmath

import "math"

type Vector2 struct {
	X, Y float32
}

func (v Vector2) Add(o Vector2) Vector2 {
	return Vector2{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vector2) Sub(o Vector2) Vector2 {
	return Vector2{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vector2) Negative() Vector2 {
	return Vector2{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vector2) Scale(s float32) Vector2 {
	return Vector2{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func (v Vector2) Dot(o Vector2) float32 {
	return v.X*o.X + v.Y*o.Y
}

func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vector2) Normalize() Vector2 {
	l := v.Length()
	if l == 0 {
		return v
	}
	return v.Scale(1 / l)
}

func (v Vector2) Perp() Vector2 {
	return Vector2{-v.Y, v.X}
}
