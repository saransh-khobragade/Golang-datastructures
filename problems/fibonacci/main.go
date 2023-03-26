package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return 1
	} else {
		result := fibo(n-1) + fibo(n-2)
		fmt.Println(result)
		return result
	}
}

func main() {
	fibo(10)
}
