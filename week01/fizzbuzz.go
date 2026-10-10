package main

import "fmt"

func fizzBuzzVariant(n int) string {
	if n%4 == 0 {
		return "Quad"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(n)
}

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println(fizzBuzzVariant(i))
	}
}
