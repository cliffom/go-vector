package vector

import (
	"math"
	"testing"
)

const epsilon = 1e-10 // tolerance for floating point equality

// Function to test float equality with tolerance
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func TestNewVector3D(t *testing.T) {
	v1 := Vector3D{3, 4, 0}
	v2 := NewVector3D(3, 4, 0)

	vSum := v1.Subtract(v2)
	vz := Vector3D{0, 0, 0}

	if vSum != vz {
		t.Errorf("v1 not equal to v2; expected zero vector but got: %v", vSum)
	}
}

func TestMagnitude(t *testing.T) {
	v := Vector3D{3, 4, 0}
	expected := 5.0
	if mag := v.Magnitude(); !almostEqual(mag, expected, epsilon) {
		t.Errorf("Magnitude() = %v, want %v", mag, expected)
	}

	// edge case: zero vector
	v = Vector3D{0, 0, 0}
	expected = 0.0
	if mag := v.Magnitude(); mag != expected {
		t.Errorf("Magnitude() of zero vector = %v, want %v", mag, expected)
	}
}

func TestNormalize(t *testing.T) {
	v := Vector3D{3, 4, 0}
	expected := Vector3D{0.6, 0.8, 0.0}
	norm, err := v.Normalize()
	if err != nil {
		t.Errorf("Normalize() returned an error: %v", err)
	}
	if !almostEqual(norm.X, expected.X, epsilon) ||
		!almostEqual(norm.Y, expected.Y, epsilon) ||
		!almostEqual(norm.Z, expected.Z, epsilon) {
		t.Errorf("Normalize() = %v, want %v", norm, expected)
	}

	// edge case: zero vector
	v = Vector3D{0, 0, 0}
	_, err = v.Normalize()
	if err == nil {
		t.Errorf("Normalize() of zero vector should return an error")
	}
}

func TestAdd(t *testing.T) {
	v1 := Vector3D{3, 4, 5}
	v2 := Vector3D{1, 2, 3}
	expected := Vector3D{4, 6, 8}
	result := v1.Add(v2)
	if result != expected {
		t.Errorf("Add() = %v, want %v", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	v1 := Vector3D{3, 4, 5}
	v2 := Vector3D{1, 2, 3}
	expected := Vector3D{2, 2, 2}
	result := v1.Subtract(v2)
	if result != expected {
		t.Errorf("Subtract() = %v, want %v", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	v := Vector3D{3, 4, 5}
	scalar := 2.0
	expected := Vector3D{6, 8, 10}
	result := v.Multiply(scalar)
	if result != expected {
		t.Errorf("Multiply() = %v, want %v", result, expected)
	}
}

func TestDotProduct(t *testing.T) {
	v1 := Vector3D{3, 4, 5}
	v2 := Vector3D{1, 2, 3}
	expected := 26.0
	if result := v1.DotProduct(v2); !almostEqual(result, expected, epsilon) {
		t.Errorf("DotProduct() = %v, want %v", result, expected)
	}
}

func TestCrossProduct(t *testing.T) {
	v1 := Vector3D{3, -3, 1}
	v2 := Vector3D{4, 9, 2}
	expected := Vector3D{-15, -2, 39}
	result := v1.CrossProduct(v2)
	if result != expected {
		t.Errorf("CrossProduct() = %v, want %v", result, expected)
	}
}

func TestDistance(t *testing.T) {
	v1 := Vector3D{1, 2, 3}
	v2 := Vector3D{4, 5, 6}
	expected := math.Sqrt(27)
	if result := v1.Distance(v2); !almostEqual(result, expected, epsilon) {
		t.Errorf("Distance() = %v, want %v", result, expected)
	}

	// edge case: same point
	v1 = Vector3D{1, 2, 3}
	v2 = Vector3D{1, 2, 3}
	expected = 0
	if result := v1.Distance(v2); result != expected {
		t.Errorf("Distance() = %v, want %v", result, expected)
	}
}
