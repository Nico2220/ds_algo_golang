package binarysearch


func LowerBound(arr []int, k int) int {
	low , hight := 0, len(arr) - 1
	ans := len(arr) - 1
	for low <= hight {
		mid := (low + hight)/ 2
		if arr[mid] >= k {
			ans = mid
			hight = mid - 1
		}else {
			low = mid + 1
		}
	}

	return ans
}