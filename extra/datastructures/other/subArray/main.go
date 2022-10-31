package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}

	for i := 1; i < len(arr)+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Println(arr[j:i])
		}
	}
}
