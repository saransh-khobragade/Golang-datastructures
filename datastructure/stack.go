package main

import "fmt"

type Node struct {
	value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.value)
}

type Stack struct {
	nodes []*Node
	top   int
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes, n)
	s.top++
}
func (s *Stack) Pop() *Node {
	if len(s.nodes) == 0 {
		return nil
	} else {
		element := s.nodes[s.top-1]
		s.nodes = append(s.nodes[:s.top-1], s.nodes[s.top:]...)
		s.top--
		return element
	}
}

func main() {
	s := &Stack{}

	s.Push(&Node{5})
	s.Push(&Node{4})
	s.Push(&Node{3})
	s.Push(&Node{2})
	s.Push(&Node{1})

	fmt.Println(s.Pop()) //1
	fmt.Println(s.Pop()) //2
	fmt.Println(s.Pop()) //3
	fmt.Println(s.Pop()) //4
	fmt.Println(s.Pop()) //5
	fmt.Println(s.Pop())
	//LIFO-Last come first serve
}
