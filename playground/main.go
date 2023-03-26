package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(str string) bool {
	l := len(str) - 1
	i := 0
	flag := true
	for i < len(str) {
		if string(str[i]) != string(str[l-i]) {
			flag = false
			break
		}
		i++
	}

	return flag
}
func main() {
	fmt.Println(strconv.Itoa(10))
	fmt.Println(isPalindrome("10"))
}
