package main

import (
	"fmt"
	"math"
)

//define interface
type Shape interface {
	area() float64
}

//define structs
type Circle struct {
	r float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{5}
	rect := Rectangle{10, 5}

	fmt.Println(getArea(rect))
	fmt.Println(getArea(circle))
}
