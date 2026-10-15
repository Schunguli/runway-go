package main

import "fmt"

func main() {
	x := 10

	if true {
		x := 20
		fmt.Println("Inside if:", x)
	}

	fmt.Println("Outside if:", x)
}
