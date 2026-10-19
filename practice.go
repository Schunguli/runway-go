package main

import "fmt"

func countPositive(numbers []int) int {
	count := 0

	for _, number := range numbers {
		if number > 0 {
			count++
		}
	}
	return count
}

func main() {
	numbers := []int{2, -1, 5, 0}
	fmt.Println(countPositive(numbers))
}
