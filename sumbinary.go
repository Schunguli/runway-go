package main

import "fmt"

func main() {
	sum := 0.1 + 0.2
	fmt.Printf("Default Println: %v\n", sum)
	fmt.Printf("Full precision: %.17f\n", sum)
}
