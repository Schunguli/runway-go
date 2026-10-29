package main

import "testing"

func TestCounts(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		letters   int
		words     int
		sentences int
	}{
		{
			name:      "one sentence",
			text:      "A small cat sleeps.",
			letters:   15,
			words:     4,
			sentences: 1,
		},
		{
			name:      "exclamation and question marks",
			text:      "Stop! Why?",
			letters:   7,
			words:     2,
			sentences: 2,
		},
		{
			name:      "extra spaces",
			text:      "  Go   is fun.  ",
			letters:   7,
			words:     3,
			sentences: 1,
		},
		{
			name:      "empty text",
			text:      "",
			letters:   0,
			words:     0,
			sentences: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := CountLetters(test.text); got != test.letters {
				t.Errorf("CountLetters() = %d, want %d", got, test.letters)
			}
			if got := CountWords(test.text); got != test.words {
				t.Errorf("CountWords() = %d, want %d", got, test.words)
			}
			if got := CountSentences(test.text); got != test.sentences {
				t.Errorf("CountSentences() = %d, want %d", got, test.sentences)
			}
		})
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
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
		{
			name: "empty text",
			text: "",
			want: "Before Grade 1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Grade(test.text); got != test.want {
				t.Errorf("Grade() = %q, want %q", got, test.want)
			}
		})
	}
}
