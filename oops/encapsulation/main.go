/*

Encapsulation means bundling everything like variables and method under one unit like class.
Go doesnt have classes instead it has structs for encapsulation.

In golang there are no access modifiers like public or private
Instead we have exported(starting with capital letter) and unexported(small letter) filed for package level
*/

package main

import "fmt"

type Customer struct {
	id   int
	name string
}

func (c *Customer) GetID() int {
	return c.id
}

func (c *Customer) GetName() string {
	return c.name
}

func main() {
	a := Customer{
		1,
		"Saransh",
	}
	fmt.Println(a.GetID())
	fmt.Println(a.GetName())
}
