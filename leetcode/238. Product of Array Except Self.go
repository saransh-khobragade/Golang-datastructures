package main

import "fmt"

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	for i := range products {
		products[i] = 1
	}

	left := 1
	for i := 1; i < len(products); i++ {
		products[i] = nums[i-1] * left
		left = products[i]
	}

	right := 1
	for i := len(products) - 1; i > -1; i-- {
		products[i] = products[i] * right
		right *= nums[i]
	}
	return products
}
func main() {
	arr := []int{2, 3, 5, 0}
	fmt.Println(productExceptSelf(arr))
}
