package main

import "fmt"

type Node struct {
	value int
}

type Stack struct {
	nodes []*Node
	top   int
	size  int
}

func (s *Stack) push(n *Node) {
	if len(s.nodes) < s.size {
		s.nodes = append(s.nodes, n)
		s.top++
	} else {
		fmt.Println("Stack overflowed")
	}
}
func (s *Stack) pop() *Node {
	if len(s.nodes) > 0 {
		ele := s.nodes[s.top-1]
		s.nodes = append(s.nodes[:s.top-1], s.nodes[s.top:]...)
		s.top--
		return ele
	} else {
		fmt.Println("Stack underflowed")
		return nil
	}
}

func main() {
	s := &Stack{
		size: 5,
	}
	s.push(&Node{5})
	s.push(&Node{4})
	s.push(&Node{3})
	s.push(&Node{2})
	s.push(&Node{1})
	s.push(&Node{6})

	fmt.Println(s.pop()) //1
	fmt.Println(s.pop()) //2
	fmt.Println(s.pop()) //3
	fmt.Println(s.pop()) //4
	fmt.Println(s.pop()) //5
	fmt.Println(s.pop())
	//LIFO-Last come first serve
}
