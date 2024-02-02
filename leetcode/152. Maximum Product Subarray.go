package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	max_sum := math.MinInt
	curr_sum := 1
	start := 0
	end := 0
	s := 0

	for i := range nums {
		curr_sum *= nums[i]
		if curr_sum > max_sum {
			max_sum = curr_sum
			start = s
			end = i
		}
		if curr_sum < 0 {
			curr_sum = 1
			s += 1
		}
	}
	fmt.Println(start, end)
	return max_sum
}

func main() {
	arr := []int{-3, -1, -1}
	fmt.Println(maxProduct(arr))
}
