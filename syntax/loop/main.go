package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}

	//Method1 for loop
	fmt.Println("--------------------")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//Method2 for each
	fmt.Println("--------------------")

	for i, v := range arr {
		fmt.Println("Value at index:", i, " is:", v)
	}

	//Method3 only value
	fmt.Println("--------------------")

	for i := range arr {
		fmt.Println("Index only", i)
	}

	//Method4 while loop
	fmt.Println("--------------------")

	j := 0
	for j < len(arr) {
		fmt.Println(arr[j])
		j++
	}

}
