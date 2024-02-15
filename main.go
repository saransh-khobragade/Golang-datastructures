package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, -1, 2}
	curr := 0
	max := math.MinInt

	start := 0
	end := 0
	s := 0

	for i := range arr {
		curr += arr[i]
		if curr > max {
			max = curr
			start = s
			end = i
		}
		if curr < 0 {
			curr = 0
			s = i + 1
		}
	}
	fmt.Println(max)
	fmt.Println(start)
	fmt.Println(end)
}
