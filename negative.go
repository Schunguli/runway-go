package main

import "fmt"

func countNegative(numbers []int) int {
	count := 0
	for _, number := range numbers {
		if number < 0 {
			count++
		}
	}
	return count
}

func main() {
	numbers := []int{2, -1, 5, -3, 0}
	fmt.Println(countNegative(numbers))
}
