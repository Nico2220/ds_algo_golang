package stringsex

import (
	"fmt"
	"strconv"
)

func LargestOddNumber(num string) string {
	// number := int(num)
	number, _ := strconv.Atoi(num)
	largest := 0
	if IsOdd(number) {
		return num
	} else {
		for _, v := range num {
			n, _ := strconv.Atoi(string(v))
			if IsOdd(n) {
				currentLargest := n
				if largest < currentLargest {
					largest = currentLargest
				}
			}
		}
	}
	if largest == 0 {
		return ""
	}
	return fmt.Sprintf("%d", largest)
}

func IsOdd(num int) bool {
	return num%2 != 0
}
