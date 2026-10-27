package main

import (
	"fmt"
	"strings"
)

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

	fmt.Printf("Player 1 score: %d\n", score1)
	fmt.Printf("Player 2 score: %d\n", score2)

	if score1 > score2 {
		fmt.Printf("Player 1 wins by %d points!\n", score1-score2)
	} else if score2 > score1 {
		fmt.Printf("Player 2 wins by %d points!\n", score2-score1)
	} else {
		fmt.Println("Tie!")
	}
}

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
