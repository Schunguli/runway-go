package main

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		person   string
		language string
		want     string
	}{
		{name: "English greeting", person: "Sasha", language: "English", want: "Hello, Sasha"},
		{name: "Spanish greeting", person: "Sasha", language: "Spanish", want: "Hola, Sasha"},
		{name: "French greeting", person: "Sasha", language: "French", want: "Bonjour, Sasha"},
		{name: "empty name defaults to World", person: "", language: "English", want: "Hello, World"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.person, tt.language)
			if got != tt.want {
				t.Errorf("Hello(%q, %q) = %q, want %q",
					tt.person, tt.language, got, tt.want)
			}
		})
	}
}
