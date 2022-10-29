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

	jim.updateByCallByValue() // Call By Value
	jim.print()               //{firstName:Jim lastName:Anderson contact:{email:abc@gmail.com zipCode:7879989}}

	jim.updateByCallByReference() //Call By Reference
	jim.print()                   //{firstName:Ankita lastName:Anderson contact:{email:abc@gmail.com zipCode:7879989}}%

}

func (p person) print() { //struct with recievers
	fmt.Printf("%+v", p)
}
func (p person) updateByCallByValue() {
	p.firstName = "Ankita"
}

func (p *person) updateByCallByReference() {
	(*p).firstName = "Ankita"
}
