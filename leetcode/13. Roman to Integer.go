// 13. Roman to Integer
// Input: s = "LVIII"
// Output: 58
package main

import "fmt"

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	out := 0
	prev := ""
	for i := len(s) - 1; i >= 0; i-- {

		curr := string(s[i])

		if prev == "" {
			out += roman[curr]
		} else if roman[curr] >= roman[prev] {
			out += roman[curr]
		} else {
			out -= roman[curr]
		}
		prev = curr
	}
	return out
}
func main() {
	s := "XLIII" //43
	fmt.Println(romanToInt(s))
}

/*
I 1
V 5
X 10
L 50
C 100
D 500
M 1000

IV 4
IX 9
XL 40
XC 90
CD 400
CM 900
*/
