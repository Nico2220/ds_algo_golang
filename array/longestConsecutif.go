package array

import (
	"math"
	"slices"
)


func LongestConsecurif(nums []int) int{
	longest := 0
	for _, v := range nums {
		nextValue := v + 1
		j := 0
		count := 1

		for j < len(nums) {
			if nums[j] == nextValue {
				count++
				nextValue++
				j = 0
			}else {
				j++
			}
		}

		if longest < count {
			longest = count
		}
	}

	return longest
}


func LongestConsecutif2(nums []int)int {
	slices.Sort(nums)
	longest := 0
	smallest := math.Inf(-1)
	count := 0
	for _, v := range nums {
		if v == int(smallest) { 
			continue
		}else if v - 1 == int(smallest) {
			count++
			smallest++
		}else {
			count = 1
			smallest = float64(v)
		}

		if longest < count {
			longest = count
		}
	}
	return longest
}

func LongestConsecutif3(nums []int) int{
	setNums := map[int]bool{}
	longest := 1
	for _, v := range nums {
		setNums[v] = true
	}
	count := 0
	smallest :=0 
	for key := range setNums {
		_, ok := setNums[key - 1]
		if !ok {
			count = 1
			smallest = key
			

			_, ok := setNums[smallest + 1] 
			for ok {
				count++
				smallest++
			}

			if longest  < count {
				longest = count
			}
		}
	}

	return longest
}


