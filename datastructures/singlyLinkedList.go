package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (list *List) insert(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}
	if list.head != nil {
		ptr := list.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = newNode
	} else {
		list.head = newNode
	}
}

func (list *List) display() {
	ptr := list.head

	for ptr.next != nil {
		fmt.Printf("%+v->", ptr.data)
		ptr = ptr.next
	}
	fmt.Printf("%+v", ptr.data)
}

func main() {
	l := List{}

	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.display()
}
