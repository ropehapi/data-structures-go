package main

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enqueue(value string) {
	node := &Node{Value: value}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func (q *Queue) Dequeue() string {
	if q.Head == nil {
		return ""
	}

	node := q.Head
	q.Head = q.Head.Next

	if q.Head == nil {
		q.Tail = nil
	}

	return node.Value
}

type Node struct {
	Value string
	Next  *Node
}
