package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("inputFile")

	if err != nil {
		log.Fatal(err)
	}

	result1 := string(content)
	fmt.Println(strings.Split(result1, "\n"))
}
