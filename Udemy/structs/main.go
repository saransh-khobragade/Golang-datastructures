package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "abc@gmail.com",
			zipCode: 7879989,
		},
	}
	jim.print()
}

func (p person) print() { //struct with recievers
	fmt.Printf("%+v", p)
}
