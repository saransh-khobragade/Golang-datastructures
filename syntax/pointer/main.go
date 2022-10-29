package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer) //0x1400000e028
	fmt.Println(*namePointer) //bill
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer) //0x1400000e038 new pointer created with different address value but referencing same old address
	fmt.Println(*namePointer) //bill
}
