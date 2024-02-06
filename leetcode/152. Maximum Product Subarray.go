package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	max := math.MinInt

	curr := 1
	for i := 0; i < len(nums); i++ {
		curr *= nums[i]
		if curr > max {
			max = curr
		}
		if curr == 0 {
			curr = 1
		}
	}
	curr = 1
	for i := len(nums) - 1; i > 0; i-- {
		curr *= nums[i]
		if curr > max {
			max = curr

		}
		if curr == 0 {
			curr = 1
		}
	}
	return max
}

func main() {
	arr := []int{-1, -2, -3, 0}
	fmt.Println(maxProduct(arr))
}
