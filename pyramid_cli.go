package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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

func PrintPyramid(height int) error {
	pyramid, err := Pyramid(height)
	if err != nil {
		return err
	}

	fmt.Print(pyramid)
	return nil
}

func heightFromArgs() (int, error) {
	if len(os.Args) != 2 {
		return 0, errors.New("usage: pyramid <height>")
	}

	height, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, errors.New("height must be a number")
	}

	return height, nil
}

func main() {
	height, err := heightFromArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if err := PrintPyramid(height); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
