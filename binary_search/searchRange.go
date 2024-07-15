package binarysearch


func SearchRange(nums []int, target int) []int {
	// result := []int{-1, -1}

	indexOne := LowboundOftarget(nums, target)
	if indexOne == len(nums){
		return []int{-1,-1}
	}
	indexTwo := UpperBoundTarget(nums, target)

	result := []int{indexOne, indexTwo -1}
	return result  
}


func LowboundOftarget(nums []int, target int) int{
	low, hight := 0 , len(nums) - 1
	ans := len(nums)

	for low <= hight  {
		mid := (low + hight) / 2
		if nums[mid] <= target {
			ans = mid
			low = mid + 1
		}else{
			hight = mid - 1
		}
	}

	return ans
}

func UpperBoundTarget(nums []int, target int) int{
	low, hight := 0 , len(nums) - 1
	ans := len(nums)

	for low <= hight  {
		mid := (low + hight) / 2
		if nums[mid] > target {
			ans = mid
			hight = mid - 1
		}else{
			low = mid + 1
		}
	}

	return ans
}