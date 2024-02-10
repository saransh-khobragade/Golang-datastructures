package main

import (
	"fmt"
	"math"
)

func trap(height []int) int {
	i := 0
	j := len(height) - 1
	maxL := height[i]
	maxR := height[j]
	out := 0

	for i < j {
		if maxL <= maxR {
			i += 1
			maxL = int(math.Max(float64(maxL), float64(height[i])))
			out += maxL - height[i]
		} else {
			j -= 1
			maxR = int(math.Max(float64(maxR), float64(height[j])))
			out += maxR - height[j]
		}
	}
	return out
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
