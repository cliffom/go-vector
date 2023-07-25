package vector

import (
	"errors"
	"math"
)

// Vector3D represents a vector in 3D space
type Vector3D struct {
	// X, Y, Z are the coordinates of a vector
	X, Y, Z float64
}

// Magnitude returns the magnitude of a vector
func (v Vector3D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize normalizes a vector into a unit vector (magnitude = 1)
func (v Vector3D) Normalize() (Vector3D, error) {
	mag := v.Magnitude()
	if mag == 0 {
		return v, errors.New("cannot normalize a zero vector")
	}

	return Vector3D{
		X: v.X / mag,
		Y: v.Y / mag,
		Z: v.Z / mag,
	}, nil
}
