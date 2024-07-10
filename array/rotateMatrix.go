package array

import (
	"fmt"
	"slices"
)

func RotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	ans := make([][]int, n)
	for i := range matrix {
		ans[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans[j][n-1-i] = matrix[i][j]
		}
	}

	return ans
}

func RotateMatrix_solution2(matrix [][]int)[][]int {
	n := len(matrix)
	for i := 0; i < n - 1; i++ {
		for j := i+1; j < n; j++ {
			swapMatrix(i, j, matrix)
		}
	}

	fmt.Println(matrix)

	for i := 0; i < n; i++ {
		slices.Reverse(matrix[i])
	}

	return matrix
}

func swapMatrix(i, j int, matrix [][]int) {
	temp := matrix[i][j]
	matrix[i][j] = matrix[j][i]
	matrix[j][i] = temp
}

func Reverse(arr []int) {
	left := 0;
	right := len(arr) - 1
	for left < right {
		tem := arr[left]
		arr[left] = arr[right]
		arr[right] = tem
		left++
		right--
	}
}
