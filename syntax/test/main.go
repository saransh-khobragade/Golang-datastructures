package main

import "fmt"

func main() {
	c := color("Red")

	fmt.Println(c)
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}
