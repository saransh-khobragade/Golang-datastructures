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
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	item, _ := s.Pop()
	fmt.Println(item)
}
