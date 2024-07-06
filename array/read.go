package array

func Read(N int, book []int, target int) string {
	seen := map[int]bool{}
	for i, v := range book {
		potential := target - v
		if seen[potential] == true {
			return "YES"
		}

		seen[book[i]] = true
	}
	return "NO"
}
