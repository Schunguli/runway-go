package main

import "testing"

func TestFromBase(t *testing.T) {
	tests := []struct {
		name string
		num  string
		base int
		want int
		ok   bool
	}{
		{"binary", "11111111", 2, 255, true},
		{"octal", "77", 8, 63, true},
		{"decimal", "255", 10, 255, true},
		{"hexadecimal", "FF", 16, 255, true},
		{"zero", "0", 10, 0, true},
		{"invalid binary digit", "102", 2, 0, false},
		{"invalid octal digit", "19", 8, 0, false},
		{"invalid hexadecimal digit", "G1", 16, 0, false},
		{"empty input", "", 10, 0, false},
		{"invalid base", "10", 1, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := FromBase(tt.num, tt.base)

			if got != tt.want || ok != tt.ok {
				t.Errorf(
					"FromBase(%q, %d) = (%d, %v), want (%d, %v)",
					tt.num, tt.base, got, ok, tt.want, tt.ok,
				)
			}
		})
	}
}

func TestToBase(t *testing.T) {
	tests := []struct {
		name string
		num  int
		base int
		want string
		ok   bool
	}{
		{"binary", 255, 2, "11111111", true},
		{"octal", 255, 8, "377", true},
		{"decimal", 255, 10, "255", true},
		{"hexadecimal", 255, 16, "FF", true},
		{"zero", 0, 2, "0", true},
		{"invalid base", 10, 1, "", false},
		{"negative number", -1, 10, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ToBase(tt.num, tt.base)

			if got != tt.want || ok != tt.ok {
				t.Errorf(
					"ToBase(%d, %d) = (%q, %v), want (%q, %v)",
					tt.num, tt.base, got, ok, tt.want, tt.ok,
				)
			}
		})
	}
}
