package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
