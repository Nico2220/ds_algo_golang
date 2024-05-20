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
	currentNode := list.head

	if currentNode == nil {
		list.head = newNode
		return
	}

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = newNode
}

func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	middleNode()

}
