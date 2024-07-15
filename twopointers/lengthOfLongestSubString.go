package twopointers

import "math"


func LenghtOfLongestSubString(s string) int{
	maxLen := 0
	for i :=0; i < len(s); i++ {
		seen := map[string]int{}
		for j :=i; j < len(s); j++ {
			v := string(s[j])
			if _, ok := seen[v]; ok {
				break;
			}

			length := j - i + 1
			maxLen = int(math.Max(float64(length), float64(maxLen)))

			seen[v] = j
		}
	}

	return maxLen
}

func LenghtOfLongestSubString2(s string) int {
	pointerOne := 0
	pointerTwo := 0
	seen := map[string]int{}
	maxLen := 0
	for pointerTwo < len(s) {
		v := string(s[pointerTwo])

		if r, ok := seen[v]; ok && r >= pointerOne  {
			pointerOne = seen[v] + 1
		}

		length := pointerTwo - pointerOne + 1 
		maxLen = int(math.Max(float64(maxLen), float64(length)))

		seen[v] = pointerTwo
		pointerTwo++
	}
	
	return maxLen

}

