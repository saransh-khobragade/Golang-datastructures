package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}
type User[T CustomData] struct {
	name string
	data T
}

func main() {
	u1 := User[string]{
		name: "",
		data: "hero",
	}
	fmt.Println(u1)

	u2 := User[int]{
		name: "",
		data: 1,
	}
	fmt.Println(u2)
}
