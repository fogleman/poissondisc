## poissondisc

Poisson Disc sampling in Go.

### Installation

    $ go get -u github.com/fogleman/poissondisc

### Documentation

https://godoc.org/github.com/fogleman/poissondisc

Learn more about Poisson-disc sampling:

https://www.jasondavies.com/poisson-disc/

https://bl.ocks.org/mbostock/dbb02448b0f93e4c82c3

### Usage

```go
var x0, y0, x1, y1, r float64
x0 = 0    // bbox min
y0 = 0    // bbox min
x1 = 1000 // bbox max
y1 = 1000 // bbox max
r = 10    // min distance between points
k := 30   // max attempts to add neighboring point

points := poissondisc.Sample(x0, y0, x1, y1, r, k, nil)

for _, p := range points {
	fmt.Println(p.X, p.Y)
}
```

### Example

    $ go run examples/example.go

![Example](https://i.imgur.com/8TJYd9s.png)
