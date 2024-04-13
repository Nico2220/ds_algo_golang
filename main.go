package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1}
	result := consecutiveOnes(arr)
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

func longestSum(arr []int, k int) int {
	start := 0
	end := 0
	sum := 0
	longest := []int{1, 2}
	for end < len(arr) {
		sum += arr[end]
		for sum > k && start < end {
			sum -= arr[start]
			start++
		}

		if sum == k && longest[1]-longest[0]+1 < end-start+1 {
			longest[0] = start
			longest[1] = end
		}

		end += 1
	}

	return longest[1] - longest[0] + 1

}
