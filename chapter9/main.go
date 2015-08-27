package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a + a*b + b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func circleArea2(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Circle struct {
	//	x float64
	//	y float64
	//	r float64
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	// new allocates memory and return a pointer
	c := new(Circle)
	// allocates memory, initializes, and returns a pointer
	c1 := Circle{x: 0, y: 0, r: 5}
	// or even a less verbose way
	c2 := Circle{0, 0, 5}

	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c1.x, c1.y, c1.r)
	fmt.Println(c2.x, c2.y, c2.r)

	fmt.Println(circleArea2(c2))

	fmt.Println(c2.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
