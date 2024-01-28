package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	max_sum := math.MinInt
	curr_sum := 0
	start := 0
	end := 0
	s := 0

	for i := range nums {
		curr_sum += nums[i]
		if curr_sum > max_sum {
			max_sum = curr_sum
			start = s
			end = i
		}
		if curr_sum < 0 {
			curr_sum = 0
			s += 1
		}
	}
	fmt.Println(start, end)
	return max_sum
}

func main() {
	arr := []int{-2, -3}
	fmt.Println(maxSubArray(arr))
}
