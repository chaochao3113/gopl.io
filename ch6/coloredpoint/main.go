package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

//Point类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

//指针接受者的方法
func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0 ,0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}
