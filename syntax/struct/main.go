package main

import "fmt"

type Person struct {
	name string
	age  int
	result
}
type result struct {
	maths   int
	physics int
}

func main() {

	sar := Person{
		name: "Saransh",
		age:  29,
		result: result{
			maths:   100,
			physics: 100,
		},
	}

	sar.age = 30 //update struct keys

	fmt.Printf("%+v", sar) //Print object  {name:Saransh age:29 result:{maths:100 physics:100}}%
}
