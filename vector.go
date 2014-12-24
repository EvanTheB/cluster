package cluster

import (
	"math"
)

var (
	ONE  = vector3{1, 1, 1}
	ZERO = vector3{0, 0, 0}
	X    = vector3{1, 0, 0}
	Y    = vector3{0, 1, 0}
	Z    = vector3{0, 0, 1}
)

type vector3 struct {
	x, y, z float64
}

func scale(a vector3, scale float64) vector3 {
	return vector3{
		a.x * scale,
		a.y * scale,
		a.z * scale,
	}
}

func (a *vector3) add(b vector3) {
	a.x += b.x
	a.y += b.y
	a.z += b.z
}

func add(a, b vector3) vector3 {
	return vector3{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
	}
}

func sub(a, b vector3) vector3 {
	return vector3{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
	}
}

func to(from, to vector3) vector3 {
	return sub(to, from)
}

func dot(a, b vector3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func cross(a, b vector3) vector3 {
	return vector3{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
	}
}

func norm(a vector3) vector3 {
	magn := mag(a)
	return vector3{
		a.x / magn,
		a.y / magn,
		a.z / magn,
	}
}

func (v *vector3) zero() {
	v.x = 0
	v.y = 0
	v.z = 0
}

func mag(a vector3) float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}
