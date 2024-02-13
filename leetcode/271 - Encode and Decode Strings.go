package main

import (
	"fmt"
	"strconv"
)

func main() {
	strs := []string{"leetcode", "code", "love", "you"}
	e := encode(strs)
	result := decode(e)
	fmt.Println(result)
}

func encode(strs []string) string {
	res := ""
	for _, x := range strs {
		l := strconv.Itoa(len(x))
		res += l
		res += "#"
		res += x
	}
	return res
}

func decode(str string) []string {
	result := []string{}
	i := 0
	l := ""
	for i < len(str) {
		if string(str[i]) != "#" {
			l += string(str[i])
			i++
		} else {
			n, _ := strconv.Atoi(l)
			s := ""
			j := i + 1
			for j <= i+n {
				s += string(str[j])
				j++
			}
			result = append(result, s)
			i = j
			l = ""
		}

	}
	return result
}
