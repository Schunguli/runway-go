package main

import (
	"fmt"
	"strings"
)

func spaces(n int) string {
	return strings.Repeat(" ", n)
}

func hashes(n int) string {
	return strings.Repeat("#", n)
}

func row(i, height int) string {
	return spaces(height-i) + hashes(i) + "\n"
}

func Pyramid(height int) string {
	if height <= 0 {
		return ""
	}

	var result strings.Builder
	for i := 1; i <= height; i++ {
		result.WriteString(row(i, height))
	}
	return result.String()
}

func main() {
	fmt.Printf("%q\n", Pyramid(3))
}
