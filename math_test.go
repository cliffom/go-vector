package vector

import (
	"testing"
)

func BenchmarkInvSqrt64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InvSqrt64(float64(n))
	}
}
