package array

func read(n int, book []int, target int) string {
	seen := make(map[int]bool, len(book))
	for _, v := range book {
		potential := target - v
		if _, ok := seen[potential]; ok {
			return "Yes"
		}
		seen[v] = true
	}
	return "No"
}
