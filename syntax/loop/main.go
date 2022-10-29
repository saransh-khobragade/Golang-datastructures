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

	//Method1 for each
	fmt.Println("--------------------")

	for i, v := range arr {
		fmt.Println("Value at index:", i, " is:", v)
	}

	//Method3 while loop
	fmt.Println("--------------------")

	j := 0
	for range arr {
		fmt.Println(arr[j])
		j++
	}

}
