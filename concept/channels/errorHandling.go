package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	errCh := make(chan error)
	wgDone := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		log.Println("First WaitGroup")
		wg.Done()
	}()
	go func() {
		log.Println("Second WaitGroup")
		errCh <- errors.New("Error on panic.com")
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(wgDone)
	}()
	select {
	case <-wgDone:
		fmt.Println("exitttttt")
		break
	case err := <-errCh:
		close(errCh)
		log.Fatal("Error encountered: ", err)
	}
	log.Println("Program worked!")
}
