package poissondisc

import (
	"math"
	"math/rand"
	"time"
)

// Sample produces points via Poisson-disc sampling.
// The points will all be within the box defined by `x0`, `y0`, `x1`, `y1`.
// No two points will be closer than the defined radius `r`.
// For each point, the algorithm will make `k` attempts to place a
// neighboring point. Increase this value for a better sampling or decrease
// it to reduce algorithm runtime.
// You may provide your own `*rand.Rand` instance or `nil` to have one
// created for you.
// Learn more about Poisson-disc sampling from the links below:
// https://www.jasondavies.com/poisson-disc/
// https://bl.ocks.org/mbostock/dbb02448b0f93e4c82c3
func Sample(x0, y0, x1, y1, r float64, k int, rnd *rand.Rand) []Point {
	if rnd == nil {
		rnd = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	}

	var result []Point
	var active []Point
	grid := newGrid(x0, y0, x1, y1, r)

	// add the first point
	{
		x := x0 + (x1-x0)/2
		y := y0 + (y1-y0)/2
		p := Point{x, y}
		grid.insert(p)
		active = append(active, p)
		result = append(result, p)
	}

	// try to add points until no more are active
	for len(active) > 0 {
		// pick a random active point
		index := rnd.Intn(len(active))
		point := active[index]
		ok := false

		// make k attempts to place a nearby point
		for i := 0; i < k; i++ {
			a := rnd.Float64() * 2 * math.Pi
			d := r * math.Sqrt(rnd.Float64() * 3 + 1)
			x := point.X + math.Cos(a)*d
			y := point.Y + math.Sin(a)*d
			if x < x0 || y < y0 || x > x1 || y > y1 {
				continue
			}
			p := Point{x, y}
			if !grid.insert(p) {
				continue
			}
			result = append(result, p)
			active = append(active, p)
			ok = true
			break
		}

		// make this point inactive if we failed to add a new point
		if !ok {
			active[index] = active[len(active)-1]
			active = active[:len(active)-1]
		}
	}

	return result
}
