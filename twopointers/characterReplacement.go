package twopointers

import "math"



func CharacterReplacement(s string, k int) int {
	maxLen := 0
	maxF := 0
	for i := 0; i < len(s); i++ {
		seen := map[byte]int{}
		for j := i; j < len(s); j++ {
			seen[s[j]]++
			
			maxF = int(math.Max(float64(maxF), float64(seen[s[j]])))

			changes := (j - i + 1) - maxF
			if changes <= k {
				maxLen = int(math.Max(float64(maxLen), float64( j - i + 1)))
			}else {
				break
			}
		}
	}

	return maxLen
}


func CharacterReplacement2(s string, k int) int {
	l :=0
	r :=0
	maxF := 0
	maxLen := 0
	seen := map[byte]int{}

	for l < len(s) && r < len(s) {
		seen[s[r]]++
		maxF = int(math.Max(float64(maxF), float64(seen[s[r]])))

		if (r - l + 1) - maxF > k {
			seen[s[l]]--
			l++
		}

		maxLen = int(math.Max(float64(maxLen), float64(r - l + 1)))
		r++
	}

	return maxLen
}