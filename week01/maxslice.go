package main

import "fmt"

func maxSlice(nums []int) int {
	if len(nums) == 0 {
		panic("empty slice")
	}
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max 
}

func main () {
	fmt.Println(maxSlice([]int{1, 5, 3, 9, 2}))
}