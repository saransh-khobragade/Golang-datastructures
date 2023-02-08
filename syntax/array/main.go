package main

import "fmt"

func main() {

	var arr = [6]int{1, 2, 3, 4}

	var arr2 = [...]string{"india", "japan", "canada"} //... will guess its length

	fmt.Println(arr) //[1 2 3 4 0 0] ,index not declared will take 0

	//At index
	fmt.Println(arr2[1]) //japan

	//SubArray
	fmt.Println(arr2[0:2]) //[india japan]

	//len and capacity
	fmt.Println(len(arr), cap(arr)) //6 6

	//Remove element at index i
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	fmt.Println(append(a[:i], a[i+1:]...)) // [A B D E]
}
