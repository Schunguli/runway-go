package main

import "fmt"

func aliasingBug() {
	data := []string{"Go", "Slices", "Are", "Fun"}
	a := data[1:3]
	b := data[2:]

	fmt.Println("Before change:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	b[0] = "Changed"

	fmt.Println("\nAfter change (bug):")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func aliasingFix() {
	data := []string{"Go", "Slices", "Are", "Fun"}
	a := data[1:3]

	b := make([]string, len(data[2:]))
	copy(b, data[2:])

	fmt.Println("\nBefore change:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	b[0] = "Changed"

	fmt.Println("\nAfter change (fixed):")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func main() {
	aliasingBug()
	aliasingFix()
}
