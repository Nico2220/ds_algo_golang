package twopointers

import "math"


func LongesSubStringWithAtMostKElement(s string, k int) int{
	maxLen := 0

	start, end := 0, 0
	seen := map[byte]int{}

	for end < len(s) {
		seen[s[end]] = end

		if len(seen) <= 2 {
			length := end- start + 1
			maxLen = int(math.Max(float64(maxLen), float64(length)))
		}

		if len(seen) > 2 {
			startIndex := seen[s[end]]
			delete(seen, s[start])
			start = startIndex + 1
		}
		end++
	}

	return maxLen
}