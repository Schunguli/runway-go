package main

import (
	"fmt"
)

func main() {
	words := []string{"hello", "héllo", "Habari 👋"}

	for _, word := range words {
		fmt.Printf("Word: %q\n", word)
		for i := 0; i < len(word); i++ {
			b := word[i]
			fmt.Printf("Byte %d: decimal=%d binary=%08b\n", i, b, b)
		}
		fmt.Printf("Total bytes: %d\n\n", len(word))
	}
}
