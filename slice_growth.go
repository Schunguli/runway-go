package main

import "fmt"

func main() {
	var s []int
	for i := 1; i <= 20; i++ {
		s = append(s, i)
		fmt.Printf("i=%2d len=%2d cap=%2d\n", i, len(s), cap(s))
	}
}
