package twopointers


func Minwindow(s string, t string) string{
	n, m:= len(s), len(t)
	l, r := 0, 0
	count := 0
	minLen := n
	startIndex := 0
	seen := map[byte]int{}

	for i:= 0; i < m; i++ {
		seen[t[i]]++
	}

	for r < n {
		v := s[r]

		if seen[v] > 0 {
			count++
		}else {
			seen[v]--
		}

		for count == m {
			if r - l + 1 < minLen {
				minLen = r - l + 1
				startIndex = l
			}

			seen[s[l]]--
			if seen[s[l]] > 0 {
				count--
			}
			l++
		}
		r++
	}

	
	
	return s[startIndex: minLen]
}