package main

import (
	"fmt"
	"os"
	"strconv"
)

func Encrypt(text string, key int) string {
	shift := rune(key % 26)
	result := []rune(text)

	for i, r := range result {
		switch {
		case r >= 'A' && r <= 'Z':
			result[i] = 'A' + (r-'A'+shift)%26
		case r >= 'a' && r <= 'z':
			result[i] = 'a' + (r-'a'+shift)%26
		}
	}

	return string(result)
}

func Decrypt(text string, key int) string {
	shift := rune(key % 26)
	result := []rune(text)

	for i, r := range result {
		switch {
		case r >= 'A' && r <= 'Z':
			result[i] = 'A' + (r-'A'-shift+26)%26
		case r >= 'a' && r <= 'z':
			result[i] = 'a' + (r-'a'-shift+26)%26
		}
	}

	return string(result)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <key>")
		os.Exit(1)
	}

	key, err := strconv.Atoi(os.Args[1])
	if err != nil || key < 0 {
		fmt.Fprintln(os.Stderr, "key must be a non-negative integer")
		os.Exit(1)
	}

	plaintext := "Hello, World!"
	ciphertext := Encrypt(plaintext, key)

	fmt.Println("Encrypted:", ciphertext)
	fmt.Println("Decrypted:", Decrypt(ciphertext, key))
}
