package main

func constructLL(arr []int) *Node {
	temp := &Node{data: 0, next: nil}
	currentNode := &Node{data: arr[0], next: nil}
	temp.next = currentNode

	for _, v := range arr {
		newNode := &Node{data: v, next: nil}
		currentNode.next = newNode
		currentNode = newNode
	}

	return temp.next
}
