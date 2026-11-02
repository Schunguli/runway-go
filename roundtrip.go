package main

import "fmt"

func Encrypt(text string, key int) string {
	key %= 26
	result := []byte(text)

	for i := 0; i < len(result); i++ {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] = byte((int(result[i]-'A')+key)%26 + 'A')
		} else if result[i] >= 'a' && result[i] <= 'z' {
			result[i] = byte((int(result[i]-'a')+key)%26 + 'a')
		}
	}

	return string(result)
}

func Decrypt(text string, key int) string {
	return Encrypt(text, 26-key%26)
}

func main() {
	tests := []struct {
		text string
		key  int
	}{
		{"HELLO", 0},
		{"HELLO", 25},
		{"HELLO", 26},
		{"ABCXYZ", 25},
		{"CaesarCipher", 3},
	}

	for _, test := range tests {
		encrypted := Encrypt(test.text, test.key)
		decrypted := Decrypt(encrypted, test.key)

		fmt.Printf("Text: %s, Key: %d\n", test.text, test.key)
		fmt.Printf("Encrypted: %s\n", encrypted)
		fmt.Printf("Decrypted: %s\n", decrypted)

		if decrypted == test.text {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
		fmt.Println()
	}
}