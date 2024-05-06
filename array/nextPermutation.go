package main

import "fmt"

func nextPermutation(A []int) []int {
	index := -1
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] < A[i+1] {
			index = i
			break
		}
	}

	if index == -1 {
		reverse(A)
		return A
	}

	for i := len(A) - 1; i >= index; i-- {
		if A[i] > A[index] {
			swap(i, index, A)
			break
		}
	}
	fmt.Println(A)

	subArray := A[index+1:]
	reverse(subArray)
	result := A[:index+1]
	result = append(result, subArray...)
	return result
	return []int{}
}

func swap(i, j int, array []int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}

func reverse(array []int) {
	left := 0
	right := len(array) - 1
	for left < right {
		swap(left, right, array)
		left++
		right--
	}
}
