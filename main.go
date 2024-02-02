package main

import "fmt"

func main() {

	arr := "ABCD"
	for i := range arr {
		fmt.Println(arr[i])
		fmt.Println(string(arr[i]))
	}
}
