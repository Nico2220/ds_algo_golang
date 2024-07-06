package array

func AlternateNumbers(a []int) []int {
	pos := 0
	neg := 1
	result := make([]int, len(a))

	for _, v := range a {
		if v < 0 {
			result[neg] = v
			neg += 2
		} else {
			result[pos] = v
			pos += 2
		}
	}
	return result
}
