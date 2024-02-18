package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	slice := []int{5, 3, 1, 2, 4}

	//Sorting
	slices.Sort(slice)
	fmt.Println(slice) //[1 2 3 4 5]

	slice = append(slice, 6)
	fmt.Println(slice) //[1 2 3 4 5 6]

	fmt.Println(len(slice)) //6

	fmt.Println(slice[1:3]) //[2 3]

}
