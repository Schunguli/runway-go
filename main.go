package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	chars := []string{"A", "a", "0", " ", "😊"}

	for _, ch := range chars {
		r := []rune(ch)[0]
		fmt.Printf("Character: %s\n", ch)
		fmt.Printf("Code point: %U\n", r)
		fmt.Printf("Decimal value: %d\n", r)
		fmt.Printf("UTF-8 byte length: %d\n\n", utf8.RuneLen(r))
	}
}
