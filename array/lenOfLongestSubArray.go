package array

import "math"


func lenOfLongestSubArray(arr []int, k int) int {
	seen := map[int]int{}
	sum := 0
	maxLen := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum == k {
			maxLen = int(math.Max(float64(maxLen), float64(i + 1)))
		}

		rem := sum - k 
		if _, ok := seen[rem]; ok {
			length := i - seen[rem]
			maxLen = int(math.Max(float64(maxLen), float64(length)))
		}

		seen[sum] = i
	}

	return maxLen
}