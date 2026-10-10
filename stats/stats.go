package main

func Stats(nums []float64) (min, maximum, mean float64) {
	if len(nums) == 0 {
		return 0, 0, 0
	}
	min, maximum = nums[0], nums[0]
	sum := 0.0 

	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > maximum {
			maximum = n 
		}
		sum += n 
	}
	mean = sum / float64(len(nums))
	return min, maximum, mean
}
