package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	lastNode := list.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = newNode
}

func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func exemple() {

	list := &LinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	list.Display()
}
