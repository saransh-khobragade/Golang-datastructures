package main

import "fmt"

type CustomMap[T comparable, V int | string] map[T]V

// comparable is any type which can be compared

func main() {
	m := make(CustomMap[int, string])
	m[3] = "3"
	fmt.Println(m)
}
