package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	var bits []byte
	for n > 0 {
		bit := n % 2
		bits = append(bits, byte('0'+bit))
		n /= 2
	}
	for i, j := 0, len(bits)-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
	return string(bits)
}

func FromBinary(s string) (int, error) {
	result := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != '0' && c != '1' {
			return 0, errors.New("invalid binary digit")
		}
		result = result*2 + int(c-'0')
	}
	return result, nil
}

func main() {
	for _, n := range []int{0, 1, 5, 73, 255} {
		my := ToBinary(n)
		std := strconv.FormatInt(int64(n), 2)
		fmt.Printf("ToBinary(%d) = %s | strconv = %s\n", n, my, std)
	}

	for _, s := range []string{"0", "1", "101", "1001001", "11111111"} {
		my, err := FromBinary(s)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		std, _ := strconv.ParseInt(s, 2, 64)
		fmt.Printf("FromBinary(%s) = %d | strconv = %d\n", s, my, std)
	}
}
