package stringsex

import (
	"math"
	"sort"
)

func LongestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	var result []byte
	n := len(strs)

	firstStr := strs[0]
	lastStr := strs[n-1]

	for i := 0; i < int(math.Min(float64(len(firstStr)), float64(len(lastStr)))); i++ {
		if firstStr[i] != lastStr[i] {
			return string(result)
		}
		result = append(result, firstStr[i])
	}

	return string(result)
}
