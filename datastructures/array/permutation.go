package main

import "fmt"

func permutation(arr []int, pos int, size int) [][]int {
	var result = make([][]int, 0)

	if pos == size {
		temp := make([]int, size)
		copy(temp, arr)
		result = append(result, temp)
	} else {
		for i := pos; i < size; i++ {
			arr[i], arr[pos] = arr[pos], arr[i]
			output := permutation(arr, pos+1, size)
			result = append(result, output...)
			arr[i], arr[pos] = arr[pos], arr[i]
		}
	}
	return result
}
func main() {
	arr := []int{1, 2, 3}
	fmt.Println(permutation(arr, 0, len(arr)))
}
