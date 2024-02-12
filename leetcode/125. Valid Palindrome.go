package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "race a car"
	fmt.Println(isPalindrome(s))
}

func clearString(str string) string {
	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	result := regex.ReplaceAllString(str, "")
	result = strings.ToLower(result)
	return result
}

func isPalindrome(s string) bool {
	str := clearString(s)
	for i := 0; i < int((len(str))/2); i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
