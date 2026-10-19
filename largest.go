package main

import "fmt"

func largestNumber(numbers []int) int {
	largest := numbers[0]

	for _, number := range numbers {
		if number > largest {
			largest = number
		}
	}
	return largest
}

func main() {
	numbers := []int{2, -1, 5, 0}
	fmt.Println(largestNumber(numbers))
}
