package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <text-file>\n", os.Args[0])
		os.Exit(2)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := strings.FieldsFunc(line, func(r rune) bool {
			return !(r >= 'a' && r <= 'z') && !(r >= '0' && r <= '9')
		})
		for _, word := range words {
			counts[word]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	words := make([]string, 0, len(counts))
	for word := range counts {
		words = append(words, word)
	}

	sort.Slice(words, func(i, j int) bool {
		if counts[words[i]] == counts[words[j]] {
			return words[i] < words[j]
		}
		return counts[words[i]] > counts[words[j]]
	})

	limit := 5
	if len(words) < limit {
		limit = len(words)
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%s: %d\n", words[i], counts[words[i]])
	}
}