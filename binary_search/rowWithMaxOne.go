package binarysearch


func RowWithMaxOne(matrix [][]int) int{
	maxCount := 0
	rowIndex := 0

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		count := 0
		for j := 0; j < len(row); j++ {
			if row[j] == 1 {
				count++
			}
		}

		if maxCount < count {
			maxCount = count
			rowIndex = i
		}
	}

	return rowIndex
}


func RowWithMaxOne2(matrix [][]int) int{
	maxCount := 0
	rowIndex := 0

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		count := 0
		
		ans := lowBound(row, 1)
		count = len(row) - ans

		if maxCount < count {
			maxCount = count
			rowIndex = i
		}
	}

	return rowIndex
}

func lowBound(arr []int, k int) int {
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