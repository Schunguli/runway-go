package main

import (
	"fmt"
	"strings"
)

func computeScore(word string, points [26]int) int {
	score := 0
	word = strings.ToUpper(word)
	for _, ch := range word {
		if ch >= 'A' && ch <= 'Z' {
			score += points[ch-'A']
		}
	}
	return score
}

func main() {
	points := [26]int{
		1, 3, 3, 2, 1, 4, 2, 4, 1, 8,
		5, 1, 3, 1, 1, 3, 10, 1, 1, 1,
		1, 4, 4, 8, 4, 10,
	}
	var word1, word2 string
	fmt.Print("Player 1: ")
	fmt.Scan(&word1)
	fmt.Print("Player 2: ")
	fmt.Scan(&word2)

	score1 := computeScore(word1, points)
	score2 := computeScore(word2, points)

	if score1 > score2 {
		fmt.Println("Player 1 wins!")
	} else if score2 > score1 {
		fmt.Println("Player 2 wins!")
	} else {
		fmt.Println("Tie!")
	}
}
