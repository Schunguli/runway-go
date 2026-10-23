package main

import (
	"errors"
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

func Pyramid(height int) (string, error) {
	if height < 1 || height > 8 {
		return "", errors.New("height must be between 1 and 8")
	}

	var result strings.Builder
	for i := 1; i <= height; i++ {
		result.WriteString(row(i, height))
	}
	return result.String(), nil
}

func main() {
	pyramid, err := Pyramid(3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print(pyramid)
}
