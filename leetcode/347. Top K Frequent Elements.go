package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}
func topKFrequent(nums []int, k int) []int {
	hmap := map[int]int{}
	for _, x := range nums {
		if _, y := hmap[x]; y {
			hmap[x] += 1
		} else {
			hmap[x] = 1
		}
	}
	temp := map[int][]int{}
	var max float64 = 0
	for x, y := range hmap {
		temp[y] = append(temp[y], x)
		max = math.Max(max, float64(y))
	}
	out := []int{}
	for i := int(max); i > 0; i-- {
		out = append(out, temp[i]...)
	}
	return out[:k]
}
