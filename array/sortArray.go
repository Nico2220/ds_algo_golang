package array

func SortdArray(arr []int, n int) []int {
	var zeros, ones, twos int

	for _, v := range arr {
		if v == 0 {
			zeros++
		} else if v == 1 {
			ones++
		} else {
			twos++
		}
	}

	k := 0
	for i := 0; i < zeros; i++ {
		arr[k] = 0
		k++
	}

	for i := 0; i < ones; i++ {
		arr[k] = 1
		k++
	}

	for i := 0; i < twos; i++ {
		arr[k] = 2
		k++
	}

	return arr
}
