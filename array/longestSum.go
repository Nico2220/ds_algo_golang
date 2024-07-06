package array

func longestSum(arr []int, k int) int {
	start := 0
	end := 0
	sum := 0
	longest := []int{1, 2}
	for end < len(arr) {
		sum += arr[end]
		for sum > k && start < end {
			sum -= arr[start]
			start++
		}

		if sum == k && longest[1]-longest[0]+1 < end-start+1 {
			longest[0] = start
			longest[1] = end
		}

		end += 1
	}

	return longest[1] - longest[0] + 1

}
