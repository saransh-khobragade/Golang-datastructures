package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readFile() {

	content, err := ioutil.ReadFile("inputFile")

	if err != nil {
		log.Fatal(err)
	}

	result := string(content)
	fmt.Println(strings.Split(result, "\n"))
}
