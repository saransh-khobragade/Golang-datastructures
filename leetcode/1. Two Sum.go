package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	hmap := map[int]int{}
	result := []int{}
	for i := range nums {
		if v, y := hmap[target-nums[i]]; y {
			result = append(result, i, v)
		} else {
			hmap[nums[i]] = i
		}
	}
	return result
}

func main() {
	input := []int{2, 7, 11, 15}
	fmt.Println(twoSum(input, 9))
}
