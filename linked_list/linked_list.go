package main

import "fmt"

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) Append(p Person) {
	node := &Node{Value: p}

	if l.Head == nil {
		l.Head = node
	}

	if l.Tail != nil {
		l.Tail.Next = node
	}

	l.Tail = node
}

func (l *List) Search(name string) (person Person) {
	node := l.Head
	for node != nil {
		if node.Value.Name == name {
			person = node.Value
			break
		}

		node = node.Next
	}

	if node != nil {
		return node.Value
	}

	return
}

func (l *List) Delete(name string) {
	if l.Head.Value.Name == name {
		l.Head = l.Head.Next
		return
	}

	prev := l.Head
	node := l.Head.Next
	for node != nil {
		if node.Value.Name == name {
			prev.Next = node.Next
			break
		}
		prev = node
		node = node.Next
	}

	if l.Tail == node {
		l.Tail = prev
	}
}

func (l *List) Display() {
	node := l.Head
	for node != nil {
		fmt.Println(node.Value.Name)
		node = node.Next
	}
}

type Node struct {
	Value Person
	Next  *Node
}
