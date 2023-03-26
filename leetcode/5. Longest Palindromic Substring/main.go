// Input: s = "babad"
// Output: "bab"
package main

import "fmt"

func main() {
	longestPalindrome("babad")
}

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
func longestPalindrome(s string) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			fmt.Println(isPalindrome(s[i : j+1]))
		}
	}
}
