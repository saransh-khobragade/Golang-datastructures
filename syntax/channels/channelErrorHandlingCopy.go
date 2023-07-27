package main

import (
	"errors"
	"log"
	"sync"
)

func main() {
	errCh := make(chan error)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		wg.Done()
	}()
	go func() {
		errCh <- errors.New("Error on panic.com")
		wg.Done()
	}()

	err := <-errCh
	close(errCh)
	log.Fatal("Error encountered: ", err)

	wg.Wait()
	log.Println("Program worked!")

}
