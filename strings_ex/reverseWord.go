package stringsex

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	sliceStrings := strings.Split(s, " ")
	cleanSlice := make([]string, 0)

	for _, v := range sliceStrings {
		if len(v) > 0 {
			cleanSlice = append(cleanSlice, v)
		}
	}

	start := 0
	end := len(cleanSlice) - 1

	for start < end {
		if len(cleanSlice[start]) == 0 {
			fmt.Println("move start")
			start++
		}
		if len(cleanSlice[end]) == 0 {
			fmt.Println("move end")
			end--
		}

		if len(cleanSlice[start]) > 0 && len(cleanSlice[end]) > 0 {
			temp := cleanSlice[start]
			cleanSlice[start] = cleanSlice[end]
			cleanSlice[end] = temp

			start++
			end--
		}

	}

	return strings.Join(cleanSlice, " ")

}
