package main

import "fmt"

func countVowels(s string) int {
	count := 0
	for _, ch := range s {
		switch ch {
			case 'a', 'e', 'i', 'o', 'u',
				'A', 'E', 'I', 'O', 'U':
				count ++	
		}
	}
	return count
}

func main() {
	fmt.Println(countVowels("hello"))
}