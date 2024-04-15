package main

import (
	"math"
)

func MaxSubArraysum(arr []int, n int) int {
	maxSum := math.Inf(-1)
	sum := 0
	for _, v := range arr {
		sum += v

		if float64(sum) > maxSum {
			maxSum = float64(sum)
		}

		if sum < 0 {
			maxSum = 0
		}
	}

	if maxSum < 0 {
		maxSum = 0
	}
	return int(maxSum)
}
