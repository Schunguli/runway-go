package main

import (
	"fmt"
	"sort"
)

func Invert(m map[string]int) map[int][]string {
	inverted := make(map[int][]string)

	for word, count := range m {
		inverted[count] = append(inverted[count], word)
	}

	for count := range inverted {
		sort.Strings(inverted[count])
	}

	return inverted
}

func main() {
	counts := map[string]int{
		"apple":  2,
		"banana": 2,
		"orange": 1,
	}

	fmt.Println(Invert(counts))
}