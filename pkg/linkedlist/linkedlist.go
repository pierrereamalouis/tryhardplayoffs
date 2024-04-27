package linkedlist

type Node struct {
	Value any   // Holds the data
	Prev  *Node // Pointer to the previous node
	Next  *Node // Pointer to the next node
}

type LinkedList struct {
	Head *Node // Pointer to the first node
	Tail *Node // Pointer to the last node
	Size int   // Tracks the size of the list
}

func (l *LinkedList) Append(value any) {
	newNode := &Node{Value: value}

	if l.Head == nil { // Empty list
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail
		l.Tail = newNode
	}
}
