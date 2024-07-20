package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	slices.Sort(nums)

	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} else {
			j := i + 1
			k := len(nums) - 1
			for j < k {
				if nums[i]+nums[j]+nums[k] < 0 {
					j++
				} else if nums[i]+nums[j]+nums[k] > 0 {
					k--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					j++
					for j < k && nums[j] == nums[j-1] {
						j++
					}
				}
			}
		}
	}
	return result
}
