package main

func ConsecutiveOnes(arr []int) int {
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
