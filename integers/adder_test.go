package integers

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		x, y int
		want int
	}{
		{name: "positive numbers", x: 2, y: 2, want: 4},
		{name: "negative and positive", x: -2, y: 5, want: 3},
		{name: "add zero", x: 7, y: 0, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
