package vector

import (
	"math"
	"testing"
)

func TestVector(t *testing.T) {
	v := &Vector3D{2, 4, 6}

	if v.X != 2 || v.Y != 4 || v.Z != 6 {
		t.Errorf("vector is %v; expected 2,4,6", v)
	}
}

func TestNormalize(t *testing.T) {
	v, _ := Vector3D{2, 4, 6}.Normalize()
	mag := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)

	if mag != 1 {
		t.Errorf("sum is %v; want 1 or less", mag)
	}
}

func TestNormalizeZeroVectir(t *testing.T) {
	_, err := Vector3D{0, 0, 0}.Normalize()

	if err == nil {
		t.Errorf("expected an error")
	}
}

func BenchmarkNormalize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v := &Vector3D{float64(n), float64(n), float64(n)}
		v.Normalize()
	}
}
