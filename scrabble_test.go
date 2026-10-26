package main

import "testing"

func TestComputeScore(t *testing.T) {
	points := [26]int{
		1, 3, 3, 2, 1, 4, 2, 4, 1, 8,
		5, 1, 3, 1, 1, 3, 10, 1, 1, 1,
		1, 4, 4, 8, 4, 10,
	}

	tests := []struct {
		word string
		want int
	}{
		{"Computer", 14},
		{"science", 11},
		{"HELLO", 8},
		{"hello", 8},
		{"world!", 9},
		{"world", 9},
	}

	for _, tt := range tests {
		got := computeScore(tt.word, points)
		if got != tt.want {
			t.Errorf("computeScore(%q) = %d, want %d", tt.word, got, tt.want)
		}
	}
}
