package main

import (
	"log"
)

func test() {
	panic("error")
}
func main() {
	defer func() {
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()

	test()
}
