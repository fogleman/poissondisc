package poissondisc

import "math"

var sentinel = Point{math.Inf(-1), math.Inf(-1)}

type grid struct {
	radius float64
	size   float64
	gw, gh int
	i0, j0 int
	cells  []Point
}

func newGrid(x0, y0, x1, y1, r float64) *grid {
	size := r / math.Sqrt(2)

	i0 := int(math.Floor(x0 / size))
	j0 := int(math.Floor(y0 / size))
	i1 := int(math.Floor(x1 / size))
	j1 := int(math.Floor(y1 / size))

	gw := i1 - i0 + 1
	gh := j1 - j0 + 1

	cells := make([]Point, gw*gh)
	for i := range cells {
		cells[i] = sentinel
	}

	return &grid{r, size, gw, gh, i0, j0, cells}
}

func (g *grid) insert(p Point) bool {
	ci := int(math.Floor(p.X/g.size)) - g.i0
	cj := int(math.Floor(p.Y/g.size)) - g.j0

	q := g.cells[cj*g.gw+ci]
	if q != sentinel && p.distance(q) < g.radius {
		return false
	}

	i0 := maxInt(0, ci-2)
	j0 := maxInt(0, cj-2)
	i1 := minInt(g.gw, ci+3)
	j1 := minInt(g.gh, cj+3)
	for j := j0; j < j1; j++ {
		for i := i0; i < i1; i++ {
			q := g.cells[j*g.gw+i]
			if q != sentinel && p.distance(q) < g.radius {
				return false
			}
		}
	}

	g.cells[cj*g.gw+ci] = p
	return true
}
