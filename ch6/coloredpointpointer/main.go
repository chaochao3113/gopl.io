package main

import (
	"fmt"
	"image/color"
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

type Rocket struct {

}

func (r *Rocket) Launch() {

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
	red := color.RGBA{255, 0 ,0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}

	fmt.Println(p.Distance(*q.Point))
	q.Point = p.Point      // p 和 q共享一个Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)

	p1 := Point{1, 2}
	q1 := Point{4, 6}

	distanceFromP := p1.Distance  // 方法变量
	fmt.Println(distanceFromP(q1))
	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p1.ScaleBy   // 方法变量
	scaleP(2)
	scaleP(3)
	scaleP(10)

	fmt.Println(p1)


	//方法变量使得程序简洁
	r := new(Rocket)
	time.AfterFunc(10 * time.Second, func() {
		r.Launch()
	})

	//简洁版
	time.AfterFunc(10 * time.Second, r.Launch)

	p2 := Point{1, 2}
	q2 := Point{4, 6}

	distance := Point.Distance  //方法表达式
	fmt.Println(distance(p2, q2))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p2, 2)
	fmt.Println(p2)
	fmt.Printf("%T\n", scale)
}