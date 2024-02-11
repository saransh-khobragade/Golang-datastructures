package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hmap := map[int]int{}

	for _, v := range nums {
		if _, y := hmap[v]; !y {
			hmap[v] = v
		}
	}
	if len(hmap) == len(nums) {
		return false
	} else {
		return true
	}
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
