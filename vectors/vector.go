package vectors

import "math"

type Vector struct {
	X int
	Y int
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (a Vector3) Subtract(b Vector3) Vector3 {
	return Vector3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func (a Vector3) Magnitude() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}
