package main

import "fmt"

// Function with int as return type
func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

//The return values of a function can be named
func area(l int, b int) (area int) {
	area = l * b
	return // Return statement without specify variable name
}

//Returning Multiple Values
func parameter(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func main() {

	sum := add(20, 30)
	fmt.Println(sum) //50

	fmt.Println("Area: ", area(20, 30)) //Area:  600

	var a, p int
	a, p = parameter(20, 30)
	fmt.Println("Area:", a)      //Area: 600
	fmt.Println("Parameter:", p) //Parameter: 100
}
