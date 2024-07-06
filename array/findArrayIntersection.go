package array

func FindArrayIntersection(arr []int, n int, brr []int, m int) []int {
	pointerOne := 0
	pointerTwo := 0
	result := []int{}
	for pointerOne < n && pointerTwo < m {
		if arr[pointerOne] < brr[pointerTwo] {
			pointerOne++
		} else if arr[pointerOne] > brr[pointerTwo] {
			pointerTwo++
		} else {
			result = append(result, arr[pointerOne])
			pointerOne++
			pointerTwo++
		}
	}
	return result
}
