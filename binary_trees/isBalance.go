package binarytrees

import "math"

func isBalance(root *Node) bool{
	return helper(root) != -1
}

func helper(root *Node)int{
	if root == nil {
		return 0
	}

	lh := helper(root.Left)
	rh := helper(root.Right)

	if lh == -1 || rh == -1 {
		return -1
	}

	if int(math.Abs(float64(lh - rh))) > 1 {
		return -1
	}

	return 1 + int(math.Max(float64(lh), float64(rh)))

}