package main

import "fmt"

type Stack []string

func (s *Stack) Push(x string) {
	*s = append(*s, x)
}
func (s *Stack) Pop() (string, bool) {
	if len(*s) < 1 {
		return "", false
	} else {
		l := len(*s) - 1
		item := (*s)[l]
		*s = (*s)[:l]
		return item, true
	}
}
func (s *Stack) Top() int {
	return len(*s)
}

func main() {
	s := "["
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var st Stack

	for _, x := range s {
		if string(x) == "(" || string(x) == "{" || string(x) == "[" {
			st.Push(string(x))
		} else if string(x) == ")" || string(x) == "}" || string(x) == "]" {
			item, _ := st.Pop()
			switch string(x) {
			case ")":
				if item != "(" {
					return false
				}
			case "}":
				if item != "{" {
					return false
				}
			case "]":
				if item != "[" {
					return false
				}
			}
		}
	}
	if st.Top() == 0 {
		return true
	} else {
		return false
	}
}
