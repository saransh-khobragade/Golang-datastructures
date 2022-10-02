package main

import (
	"fmt"
	"math"
)

func isPrime(n float64) bool {
	flag := true

	if n > 1 {
		for i := n - 1; i >= math.Sqrt(n); i-- {
			if math.Mod(n, i) == 0 {
				flag = false
				break
			}
		}
	} else {
		flag = false
	}

	return flag
}

func main() {

	n := 100
	for i := 0; i < n; i++ {
		if isPrime(float64(i)) {
			fmt.Println(i, "is Prime")
		}
	}
}

//	check number n-1 to its square root of n
//	number should not divisible by anyone else
//	number should be greater than 1
