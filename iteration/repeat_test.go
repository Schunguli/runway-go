package iteration

import "testing"

func TestRepeat(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		repeat int
		want   string
	}{
		{name: "repeat a letter five times", input: "a", repeat: 5, want: "aaaaa"},
		{name: "repeat a word three times", input: "Go", repeat: 3, want: "GoGoGo"},
		{name: "repeat zero times", input: "x", repeat: 0, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Repeat(tt.input, tt.repeat)
			if got != tt.want {
				t.Errorf("Repeat(%q, %d) = %q, want %q",
					tt.input, tt.repeat, got, tt.want)
			}
		})
	}
}
