package main

import "fmt"

func main() {
	fmt.Println(strToArray("aas"))
}

func strToArray(str string) []string {
	char := []string{}
	for _, v := range str {
		char = append(char, string(v))
	}
	return char
}
