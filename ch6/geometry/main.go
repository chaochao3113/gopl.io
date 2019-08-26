package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

//Path 是连接多个点的直线段
type Path []Point

//普通的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

//Point类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

//指针接受者的方法
func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

// nil是一个合法的接受者

// IntList是整型链表
// *IntList的类型nil代表空列表
type IntList struct {
	Value int
	Tail *IntList
}

// Sum返回列表元素的总和
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) //函数调用
	fmt.Println(p.Distance(q))  //方法调用

	//构建三角形
	perim := Path{
		{1, 1},  //起点
		{5, 1},  //第二个点
		{5, 4},  //第三个点
		{1, 1},  //终点
	}
	//输出周长
	fmt.Println(perim.Distance())

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p = Point{1, 2}
	pptr := &p
	pptr.ScaleBy(20)
	fmt.Println(*pptr)

	p = Point{1, 2}
	(&p).ScaleBy(23)
	fmt.Println(p)

	//编译器会对变量进行&p的隐式转换,只有变量才允许这么做,包括结构体字段、像p.x和数组或者slice的元素,比如perim[0]
	p.ScaleBy(2)  // 隐式转换为(&p)
	fmt.Println(p)

	//以下都是合法的
	Point{1, 2}.Distance(q)  // Point
	pptr.ScaleBy(2)  // *Point

}
