package main

import "fmt"

func main() {

	//Native print
	println("hello")

	arr := []int{10, 20, 30}

	//Raw print
	fmt.Print(arr)

	// add \n at end
	fmt.Println(arr)

	// can use formatting inline variables
	fmt.Printf("Value of variable is %s ", "abc") //string
	fmt.Printf("Value of variable is %d ", 789)   //int
	fmt.Printf("Value of variable is %T ", 789)   //type
}
