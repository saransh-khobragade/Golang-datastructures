package main

import "fmt"

func main() {
	arr := customType{1, 2, 3, 4, 5}

	for i, value := range arr {
		fmt.Println(i, value)
	}
}

//wont work if your workspace have multiple main, open olny folder which have single main.go
