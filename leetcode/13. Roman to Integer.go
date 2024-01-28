// 13. Roman to Integer
// Input: s = "LVIII"
// Output: 58
package main

import "fmt"

func main() {
	fmt.Println(romanToInt("I"))
}

func romanToInt(s string) int {
	i := 0
	result := 0
	for i < len(s) {
		switch string(s[i]) {
		case "I":
			{
				if i+1 < len(s) && string(s[i+1]) == "V" {
					result += 4
					i++
				} else if i+1 < len(s) && string(s[i+1]) == "X" {
					result += 9
					i++
				} else {
					result += 1
				}
			}
		case "V":
			result += 5
		case "X":
			{
				if i+1 < len(s) && string(s[i+1]) == "L" {
					result += 40
					i++
				} else if i+1 < len(s) && string(s[i+1]) == "C" {
					result += 90
					i++
				} else {
					result += 10
				}
			}

		case "L":
			result += 50
		case "C":
			{
				if i+1 < len(s) && string(s[i+1]) == "D" {
					result += 400
					i++
				} else if i+1 < len(s) && string(s[i+1]) == "M" {
					result += 900
					i++
				} else {
					result += 100
				}
			}
		case "D":
			result += 500
		case "M":
			result += 1000
		}
		i++
	}
	return result
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
