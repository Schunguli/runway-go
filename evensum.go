package main

import "fmt"

func sumEven(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		if number%2 == 0 {
			sum += number
		}
	}
	return sum
}

func main() {
	numbers := []int{2, 3, 4, 5}
	fmt.Println(sumEven(numbers))
}
