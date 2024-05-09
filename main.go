package main

import (
	"fmt"
)

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	result := arr[1:4]
	result = append(result, "z")
	fmt.Println(result)
	fmt.Println(arr)

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
