package main

import "fmt"

// Node represents an element in the list
type Node struct {
	data int
	next *Node
}

// LinkedList manages the nodes
type LinkedList struct {
	head *Node
}

// Prepend adds a node to the beginning of the list (O(1))
func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
}

func LinkedListTest() {
	myList := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}

	myList.Prepend(node1)
	myList.Prepend(node2)

	// Traverse the list
	curr := myList.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}
