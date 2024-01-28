// Input: s = "babad"
// Output: "bab"
package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("aabba"))
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
func longestPalindrome(s string) string {
	maxCount := 0
	output := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			l := len(s[i : j+1])
			if maxCount < l && isPalindrome(s[i:j+1]) {
				if l > maxCount {
					output = s[i : j+1]
					maxCount = l
				}
			}
		}
	}
	return output
}
