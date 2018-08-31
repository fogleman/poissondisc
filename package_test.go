package poissondisc

import (
	"math"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	size := math.Sqrt(float64(b.N) * 1.618)
	Sample(0, 0, size, size, 1, 16, nil)
}
