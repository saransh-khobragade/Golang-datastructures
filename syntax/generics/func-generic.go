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
	result := MapValues([]float64{1.1, 2.2, 3.3}, func(n float64) float64 {
		return n * 2
	})
	fmt.Println(result)
}
