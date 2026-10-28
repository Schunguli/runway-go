package main

import "fmt"

func Insert(s []int, i, v int) []int {
	if i < 0 || i > len(s) {
		return append([]int(nil), s...)
	}
	result := make([]int, len(s)+1)
	copy(result, s[:i])
	result[i] = v
	copy(result[i+1:], s[i:])
	return result
}

func RemoveAt(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return append([]int(nil), s...)
	}
	result := make([]int, len(s)-1)
	copy(result, s[:i])
	copy(result[i:], s[i+1:])
	return result
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Insert start:", Insert(s, 0, 99))
	fmt.Println("Insert middle:", Insert(s, 2, 99))
	fmt.Println("Insert end:", Insert(s, len(s), 99))
	fmt.Println("Insert out-of-range:", Insert(s, 10, 99))
	fmt.Println("Remove start:", RemoveAt(s, 0))
	fmt.Println("Remove middle:", RemoveAt(s, 2))
	fmt.Println("Remove end:", RemoveAt(s, len(s)-1))
	fmt.Println("Remove out-of-range:", RemoveAt(s, 10))
	fmt.Println("Original unchanged:", s)
}
