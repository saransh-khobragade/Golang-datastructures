package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan<- int, wg *sync.WaitGroup) { //send only channel
		ch <- 1
		ch <- 2
		ch <- 3
		close(myCh) //only sender should close
		wg.Done()
	}(myCh, wg)
	go func(ch <-chan int, wg *sync.WaitGroup) { //recieve only channel
		for v := range myCh {
			fmt.Println(v)
		}
		wg.Done()
	}(myCh, wg)
	wg.Wait()

}
