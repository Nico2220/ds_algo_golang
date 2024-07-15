package array

func MergeArray(arr1, arr2 []int) []int {
	
	p1, p2 := 0, 0
	result := []int{}

	for p1 < len(arr1) || p2 < len(arr2){
		if (arr1[p1] < arr2[p2] && p1 < len(arr1)) || p2 >= len(arr2) {
			result = append(result, arr1[p1])
			p1++
		}else {
			result = append(result, arr2[p2]) 
			p2++
		}
	}

	return result
}