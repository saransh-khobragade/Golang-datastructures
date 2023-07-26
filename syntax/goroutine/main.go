package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func ping(link string) {
	defer wg.Done() // will executue whatever happens at the end, and will give closure to every wait group menber as done

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(link, "is up")
}
func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://gsmareana.com",
	}

	for _, link := range links {
		wg.Add(1)     // adding number of threads it has to watch
		go ping(link) //making function async
	}
	wg.Wait() //will hold main function till wait group member are done
}
