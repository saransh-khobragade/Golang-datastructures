package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var maxArea float64 = 0
	for i < j {
		maxArea = math.Max(maxArea, math.Min(float64(height[i]), float64(height[j]))*float64(j-i))

		if height[i] <= height[j] {
			i += 1
		} else {
			j -= 1
		}
	}
	return int(maxArea)
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
