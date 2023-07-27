package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {

	myCh := make(chan int)
	errCh := make(chan error)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		for v := range ch {
			fmt.Println(v)
		}
		for v := range errCh {
			fmt.Println(v)
		}
		err := <-errCh
		close(errCh)

		log.Fatal("Error encountered: ", err)
		defer wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		errCh <- errors.New("Error on panic.com")
		defer close(myCh)
		defer close(errCh)
		defer wg.Done()
	}(myCh, wg)

	wg.Wait()
}
