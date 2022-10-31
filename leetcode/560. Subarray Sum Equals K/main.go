package main

import (
	"fmt"
)

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func subarraySum(nums []int, k int) int {
	count := 0
	for i := 1; i < len(nums)+1; i++ {
		for j := 0; j < i; j++ {
			if k == sum(nums[j:i]) {
				count++
			}
		}
	}
	return count
}

func main() {

	input1, input2 := readFile("inputFile")

	fmt.Println(subarraySum(input1, input2[0]))
}
