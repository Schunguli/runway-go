package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func CountLetters(text string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		char := text[i]
		if (char >= 'A' && char <= 'Z') ||
			(char >= 'a' && char <= 'z') {
			count++
		}
	}
	return count
}

func CountWords(text string) int {
	return len(strings.Fields(text))
}

func CountSentences(text string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '.', '!', '?':
			count++
		}
	}
	return count
}

func Grade(text string) string {
	words := CountWords(text)
	if words == 0 {
		return "Before Grade 1"
	}

	letters := CountLetters(text)
	sentences := CountSentences(text)

	L := float64(letters) / float64(words) * 100
	S := float64(sentences) / float64(words) * 100
	index := 0.0588*L - 0.296*S - 15.8
	grade := int(math.Round(index))

	switch {
	case grade < 1:
		return "Before Grade 1"
	case grade >= 16:
		return "Grade 16+"
	default:
		return fmt.Sprintf("Grade %d", grade)
	}
}

func main() {
	text, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read text:", err)
		os.Exit(1)
	}

	fmt.Println(Grade(string(text)))
}

func runTableTests() bool {
	countTests := []struct {
		name          string
		text          string
		wantLetters   int
		wantWords     int
		wantSentences int
	}{
		{
			name:          "counts letters, words, and sentences",
			text:          "Hello, world! It's 2026.",
			wantLetters:   13,
			wantWords:     4,
			wantSentences: 2,
		},
		{
			name:          "handles empty text",
			text:          "",
			wantLetters:   0,
			wantWords:     0,
			wantSentences: 0,
		},
	}

	gradeTests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "before Grade 1",
			text: "One fish. Two fish. Red fish. Blue fish.",
			want: "Before Grade 1",
		},
		{
			name: "Grade 2",
			text: "Would you like them here or there? I would not like them here or there. I would not like them anywhere.",
			want: "Grade 2",
		},
		{
			name: "Grade 3",
			text: "Congratulations! Today is your day. You're off to Great Places! You're off and away!",
			want: "Grade 3",
		},
		{
			name: "Grade 7",
			text: "In my younger and more vulnerable years my father gave me some advice that I've been turning over in my mind ever since.",
			want: "Grade 7",
		},
		{
			name: "Grade 16 or higher",
			text: "A large class of computational problems involve the determination of properties of graphs, digraphs, integers, arrays of integers, finite families of finite sets, boolean formulas and elements of other countable domains.",
			want: "Grade 16+",
		},
	}

	allPassed := true

	for _, test := range countTests {
		gotLetters := CountLetters(test.text)
		gotWords := CountWords(test.text)
		gotSentences := CountSentences(test.text)

		passed := gotLetters == test.wantLetters &&
			gotWords == test.wantWords &&
			gotSentences == test.wantSentences

		if passed {
			fmt.Println("PASS:", test.name)
		} else {
			allPassed = false
			fmt.Println("FAIL:", test.name)
			fmt.Printf("  Letters: got %d, want %d\n", gotLetters, test.wantLetters)
			fmt.Printf("  Words: got %d, want %d\n", gotWords, test.wantWords)
			fmt.Printf("  Sentences: got %d, want %d\n", gotSentences, test.wantSentences)
		}
	}

	for _, test := range gradeTests {
		got := Grade(test.text)
		if got == test.want {
			fmt.Println("PASS:", test.name)
		} else {
			allPassed = false
			fmt.Printf("FAIL: %s: got %q, want %q\n", test.name, got, test.want)
		}
	}

	return allPassed
}
