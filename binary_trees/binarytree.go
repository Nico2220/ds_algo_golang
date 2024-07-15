package binarytrees

type Node struct {
	Value int
	Left *Node
	Right *Node
}

// func (n *Node) Insert(value int){

// }

func(n  *Node) InOrderTraverse()[]int{
	result := []int{}
	if n == nil {
		return nil
	}

	n.Left.InOrderTraverse()
	result = append(result, n.Value)
	n.Right.InOrderTraverse()
	
	return result
}

func PreorderTraversal(root *Node) []int {
    result := []int{}

	if root == nil {
        return []int{}
    }

	result = append(result, root.Value)
    result = append(result, PreorderTraversal(root.Left)...)
	result = append(result, PreorderTraversal(root.Right)...)

	return result
}



//root := binarytrees.Node{Value: 1}
	// root.Left = &binarytrees.Node{Value:2 }
	// root.Right = &binarytrees.Node{Value: 3}

	// root.Left.Left = &binarytrees.Node{Value: 4}
	// root.Left.Right = &binarytrees.Node{Value: 5}

	// root.Right.Left = &binarytrees.Node{Value: 6}
	// root.Right.Right = &binarytrees.Node{Value: 7}