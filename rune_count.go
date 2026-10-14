package main

import (
	"fmt"
	"unicode/utf8"
)

func RuneCount(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func main() {
	words := []string{"Kisumu", "héllo", "Habari 👋"}

	for _, w := range words {
		fmt.Printf("Word: %q\n", w)
		fmt.Printf("len(s) = %d (bytes)\n", len(w))
		fmt.Printf("RuneCount(s) = %d (runes)\n", RuneCount(w))
		fmt.Printf("utf8.RuneCountInString(s) = %d (library check)\n\n",
			utf8.RuneCountInString(w))
	}
}
