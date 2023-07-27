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

// NewVector returns a new Vector3D
func NewVector3D(x, y, z float64) Vector3D {
	return Vector3D{X: x, Y: y, Z: z}
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

// Add adds two vectors together
func (v Vector3D) Add(w Vector3D) Vector3D {
	return Vector3D{
		X: v.X + w.X,
		Y: v.Y + w.Y,
		Z: v.Z + w.Z,
	}
}

// Subtract subtracts one vector from another
func (v Vector3D) Subtract(w Vector3D) Vector3D {
	return Vector3D{
		X: v.X - w.X,
		Y: v.Y - w.Y,
		Z: v.Z - w.Z,
	}
}

// Multiply scales a vector by a scalar
func (v Vector3D) Multiply(scalar float64) Vector3D {
	return Vector3D{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

// DotProduct calculates the dot product of two vectors
func (v Vector3D) DotProduct(w Vector3D) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// CrossProduct calculates the cross product of two vectors
func (v Vector3D) CrossProduct(w Vector3D) Vector3D {
	return Vector3D{
		X: v.Y*w.Z - v.Z*w.Y,
		Y: v.Z*w.X - v.X*w.Z,
		Z: v.X*w.Y - v.Y*w.X,
	}
}

// Distance calculates the distance between two vectors
func (v Vector3D) Distance(w Vector3D) float64 {
	difference := v.Subtract(w)
	return difference.Magnitude()
}
