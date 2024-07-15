package twopointers

import "math"


func NumberOfSubstrings(s string) int{
	numberOfString := 0
	
	for i := 0; i < len(s); i++ {
		seen := map[byte]bool{}
		for j := i; j < len(s); j++ {
			seen[s[j]] = true
			if len(seen) == 3 {
				numberOfString += j - i 
			}

		}


	}

	return numberOfString
}


func NumberOfSubstrings2(s string)int {
	lastseen := []int{-1, -1, -1}
	count := 0

	for i :=0; i < len(s); i++ {
		v := s[i]
		lastseen[v -  'a'] = i
		if lastseen[0] != -1 && lastseen[1] != -1 && lastseen[2] != -1 {
			min := int(math.Min(float64(lastseen[0]), float64(lastseen[1])))
			min = int(math.Min(float64(min), float64(lastseen[2])))
			count += 1 + min
		}
	}

	return count
}




