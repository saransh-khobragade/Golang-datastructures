/*
The capability of a class to derive properties and characteristics from another class is called Inheritance
Inheritance achieved by composition by embedding structs
*/
package main

import "fmt"

type Vehicle struct {
	Seats int
	Color string
}

type Car struct {
	Vehicle //composition
}

type MotorCycle struct {
	Base Vehicle //composition with attribute
}

func main() {
	//Direct access
	car := &Car{
		Vehicle{
			Seats: 4,
			Color: "blue",
		},
	}

	fmt.Println(car.Seats)
	fmt.Println(car.Color)

	//Accessed via some attribute
	motorCycle := &MotorCycle{
		Vehicle{
			Seats: 2,
			Color: "red",
		},
	}

	fmt.Println(motorCycle.Base.Seats)
	fmt.Println(motorCycle.Base.Color)
}
