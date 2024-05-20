package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	currentNode := head
	count := 0
	for currentNode != nil {
		currentNode = currentNode.Next
		count++
	}

	middle := int(count / 2)
	currentNode = head
	for i := 0; i < middle; i++ {
		currentNode = currentNode.Next
	}

	return currentNode
}
