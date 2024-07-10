package main

import (
	"fmt"

	"github.com/Nico2220/ds_algo_golang/array"
)

type Parent struct{}

func (c *Parent) Print() {
	fmt.Println("parent")
}

type Child struct {
	Parent
}

func (p *Child) Print() {
	fmt.Println("child")
}

func main() {
	result := array.RotateMatrix_solution2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(result)

}

func findArrayIntersection(arr []int, n int, brr []int, m int) []int {
	pointerOne := 0
	pointerTwo := 0
	result := []int{}
	for pointerOne < n && pointerTwo < m {
		if arr[pointerOne] < brr[pointerTwo] {
			pointerOne++
		} else if arr[pointerOne] > brr[pointerTwo] {
			pointerTwo++
		} else {
			result = append(result, arr[pointerOne])
			pointerOne++
			pointerTwo++
		}
	}
	return result
}

func consecutiveOnes(arr []int) int {
	longest := 0
	currCount := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			currCount++
			if longest < currCount {
				longest = currCount
			}
		} else {
			currCount = 0
		}
	}
	return longest
}
