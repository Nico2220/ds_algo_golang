package array

func MajorityElement(arr []int) int {
	pointerOne := 0
	pointerTwo := 0
	count := 0

	for pointerOne < len(arr) && pointerTwo < len(arr) {

		if arr[pointerOne] == arr[pointerTwo] {
			count++
		} else {
			count--
		}

		if count == 0 {
			pointerOne = pointerTwo + 1
		}

		pointerTwo++
	}

	countTwo := 0
	if count > 0 {
		for _, v := range arr {
			if v == arr[pointerOne] {
				countTwo++
			}

			if countTwo == int(len(arr)/2) {
				return arr[pointerOne]
			}
		}
	}

	return -1
}
