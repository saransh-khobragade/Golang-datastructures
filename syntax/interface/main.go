package main

import "fmt"

type person1 struct{}
type person2 struct{}

type person interface {
	getName() string //just providing abstraction name of function not body
}

func (person1) getName() string {
	return "Saransh"
}

func (person2) getName() string {
	return "Ankita"
}
func printName(p person) {
	fmt.Println(p.getName()) //from interface it just assume it would have some method getName with some implementation
}

func main() {
	p1 := person1{}
	p2 := person2{}

	printName(p1)
	printName(p2)
}

//because person1 and person2 both has a method getName they belong to person interface now
//printName can be called by both person1 and person2 instance because its knows instance will have getName method.
