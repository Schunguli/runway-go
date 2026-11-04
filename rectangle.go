package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Scale(factor float64) Rectangle {
	return Rectangle{
		Width:  r.Width * factor,
		Height: r.Height * factor,
	}
}

func main() {
	r1 := Rectangle{
		Width:  10,
		Height: 5,
	}

	r2 := r1.Scale(2)

	fmt.Println(r1)
	fmt.Println(r2)
}