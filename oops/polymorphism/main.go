/*
Polymorphism as the ability of a message to be displayed in more than one form
Objects of different types are treated in a consistent way, as long as they stick to a single interface, which is the essence of polymorphism*/

package main

import (
	"fmt"
)

type Figure interface {
	Area() float64
}

type Rectangle struct {
	length float64
	width  float64
}
type Square struct {
	side float64
}

func (rect Rectangle) Area() float64 {
	area := rect.length * rect.width
	return area
}

func (sq Square) Area() float64 {
	area := sq.side * sq.side
	return area
}

func main() {
	rectangle := Rectangle{
		length: 10.5,
		width:  12.25,
	}

	square := Square{
		side: 15.0,
	}

	// The Figure interface can hold rectangle and square type as they both implements the interface
	var f1 Figure = rectangle
	var f2 Figure = square
	// Variable declared in an interface are of interface type. They can take whichever value implements the interface which helps interfaces to achieve polymorphism in the Golang

	// printing the calculated result
	fmt.Printf("Area of rectangle: %.3f unit sq.\n", f1.Area())
	fmt.Printf("Area of square: %.3f unit sq.\n", f2.Area())
}
