package main

import (
	"errors"
	"fmt"
)

func DivMod(a, b int) (quotient, remainder int, err error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	quotient = a / b
	remainder = a % b
	return quotient, remainder, nil
}

func main() {
	q, r, err := DivMod(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
	}
}
