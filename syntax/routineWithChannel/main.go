package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://gsmareana.com",
	}

	c := make(chan string) //channel for cross routine communication

	for _, link := range links {
		go call(link, c) //making function async
	}

	for l := range c { //reading from channels(child routines) whenever channel will have value it will assign to l
		go func(l string) {
			time.Sleep(3 * time.Second) //Sleeping routine
			go call(l, c)               //you have recieve channel so that parent routine wont exit
		}(l) //passing to function literal like anonymous function
	}

}

func call(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- link //sending back to parent routine,program will hang if no one listing
		return
	}

	fmt.Println(link, "is up")
	c <- link //sending back to parent routine,program will hang if no one listing
}
