package poissondisc

import "math"

type Point struct {
	X, Y float64
}

func (p Point) distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
