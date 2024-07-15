package binarytrees

import "math"



func maxDepth(root *Node) int{
	if root == nil {
		return 0
	}

	dl := maxDepth(root.Left)
	dr := maxDepth(root.Right)

	return 1 + int(math.Max(float64(dl), float64(dr)))
}