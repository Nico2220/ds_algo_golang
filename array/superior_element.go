package main

import "slices"

func superiorElement(a []int) []int {
	max := 0
	result := []int{}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] > max {
			max = a[i]
			result = append(result, max)
		}
	}

	slices.Sort(result)
	return result
}
