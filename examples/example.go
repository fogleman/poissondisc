package main

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/fogleman/poissondisc"
)

const (
	W = 2400
	H = 1600
	R = 8
	K = 32
)

func main() {
	points := poissondisc.Sample(0, 0, W, H, R, K, nil)
	fmt.Println(len(points), "points")

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	for _, p := range points {
		dc.DrawPoint(p.X, p.Y, R*0.45)
	}
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}
