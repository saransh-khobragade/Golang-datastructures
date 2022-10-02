package main

import (
	"fmt"
)

func main() {
	var i = 5

	// AND operator
	if i > 1 && i < 10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// OR operator
	if i > 6 || i < 10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	//EQUAL and else if
	if i == 9 {
		fmt.Println("9")
	} else if i == 5 {
		fmt.Println("elseIf")
	} else {
		fmt.Println("else")
	}

}
