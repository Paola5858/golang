package main

import "math"

type Vector struct {
	X, Y float64
}

func (v *Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() {
	l := v.Len()
	if l == 0 {
		return
	}
	v.X /= l
	v.Y /= l
}

func (v Vector) Scaled(s float64) Vector {
	return Vector{v.X * s, v.Y * s}
}
