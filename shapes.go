package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	rectangle := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 10}

	fmt.Println("Rectangle Area:", rectangle.Area())
	fmt.Println("Rectangle Perimeter:", rectangle.Perimeter())

	fmt.Println("Circle Area:", circle.Area())
	fmt.Println("Circle Perimeter:", circle.Perimeter())
}