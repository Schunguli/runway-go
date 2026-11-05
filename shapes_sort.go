package main

import (
	"fmt"
	"math"
	"sort"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func SortShapesByArea(shapes []Shape) {
	sort.Slice(shapes, func(i, j int) bool {
		return shapes[i].Area() < shapes[j].Area()
	})
}

func main() {
	shapes := []Shape{
		Rectangle{4, 5}, // 20
		Triangle{4, 3},  // 6
		Circle{1},       // 3.14
	}

	SortShapesByArea(shapes)

	for _, s := range shapes {
		fmt.Printf("%.2f\n", s.Area())
	}
}