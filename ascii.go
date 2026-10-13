package main

import (
	"fmt"
)

func IsASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

func main() {
	tests := []string{"hello", "héllo", "Habari 👋"}
	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, IsASCII(t))
	}
}
