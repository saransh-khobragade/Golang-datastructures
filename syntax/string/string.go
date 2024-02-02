package main

import "fmt"

func main() {

	arr := "ABCD"
	for i := range arr {
		fmt.Println(arr[i])         //65 66 67 68
		fmt.Println(string(arr[i])) //A B C D
	}

	for i := range arr {
		if arr[i] == 'B' {
			fmt.Println("B is there")
		}
	}
}
