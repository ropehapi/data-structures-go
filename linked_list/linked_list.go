package main

import "fmt"

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) Prepend(p Person) {
	node := &Node{Value: p, Next: l.Head}
	l.Head = node
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

func (l *List) InsertAtGivenPosition(p Person, position int) {
	node := &Node{Value: p}

	if position == 1 {
		node.Next = l.Head
		l.Head = node
		return
	}

	previous := l.Head
	count := 1

	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	node.Next = current
	previous.Next = node
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

func (l *List) DeleteAtGivenPosition(position int) {
	if position <= 0 {
		return
	}

	if position == 1 {
		l.Head = l.Head.Next
		return
	}

	count := 1
	previous := l.Head

	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	previous.Next = current.Next
	current = nil
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
