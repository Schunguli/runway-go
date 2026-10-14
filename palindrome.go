package main

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	var filtered []rune
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			filtered = append(filtered, unicode.ToLower(r))
		}
	}
	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}
	return true
}

func main() {
	tests := []string{
		"Kisumu",
		"héllo",
		"Habari 👋",
		"A man, a plan, a canal: Panama!",
		"No lemon, no melon",
	}

	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, IsPalindrome(t))
	}
}
