package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(ch chan int, w *sync.WaitGroup, i int) {
			myCh <- i
			wg.Done()
		}(myCh, wg, i)
	}

	go func() {
		wg.Wait()
		close(myCh) //both sender cant close,unless other wont be able to send so used one more routine for waiting above two done and then close
	}()
	for v := range myCh {
		fmt.Println(v)
	}

}
