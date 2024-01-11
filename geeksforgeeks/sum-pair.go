//https://www.geeksforgeeks.org/check-if-pair-with-given-sum-exists-in-array/

package main

import "fmt"

func getPair(arr []int, sum int) {
	hashMap := make(map[int]int)

	for _, x := range arr {
		y := sum - x
		if _, ok := hashMap[y]; ok {
			fmt.Println(y, x)
			return
		}
		hashMap[x] = 1
	}
}

func main() {
	arr := []int{1, 4, 45, 6, 10, 8}
	N := 16
	getPair(arr, N)
}
