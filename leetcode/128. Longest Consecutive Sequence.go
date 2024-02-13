package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{4, 2, 2, -4, 0, -2, 4, -3, -4, -4, -5, 1, 4, -9, 5, 0, 6, -8, -1, -3, 6, 5, -8, -1, -5, -1, 2, -9, 1}
	fmt.Println(longestConsecutive(nums))
}
func longestConsecutive(nums []int) int {
	hmap := map[int][]int{}
	visited := map[int]bool{}

	var max float64 = 0

	for _, x := range nums {
		if _, y := visited[x]; !y {
			visited[x] = true

			LeftMost := x
			RightMost := x

			if _, y := hmap[x+1]; y {
				RightMost = hmap[x+1][1]
			}
			if _, y := hmap[x-1]; y {
				LeftMost = hmap[x-1][0]
			}
			hmap[LeftMost] = []int{LeftMost, RightMost}
			hmap[RightMost] = []int{LeftMost, RightMost}
			max = math.Max(max, float64(RightMost-LeftMost)+1)
		}

	}
	return int(max)
}
