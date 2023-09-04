// Composition by embedding structs not inheritance
package main

import "fmt"

type Vehicle struct {
	ModelNo        string
	YearOfPurchase string
}

// Method belong to Vehicle class
func (v *Vehicle) VehicleDetails() {
	fmt.Println(v.ModelNo)
	fmt.Println(v.YearOfPurchase)
}

type Car struct {
	Vehicle //embedded variable
	Company string
}

// Method belong to Car class
func (v *Car) CarCompany() {
	fmt.Println(v.Company)
}

func main() {
	c := Car{
		Vehicle{
			ModelNo:        "123",
			YearOfPurchase: "2019",
		},
		"ferrari",
	}
	c.CarCompany()     //Car'object able to access Car's method
	c.VehicleDetails() //Car'object able to access Vehicle's method
}
