package main

import (
	"fmt"
	"time"
)

func f(from string, c chan int) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(from, ":", i)
		c <- i
	}
}

func main() {

	c := make(chan int)

	go f("goroutine", c)

	for l := range c { //reading from channels(child routines) whenever channel will have value it will assign to l
		fmt.Println(l)
	}
}
