package stringsex

func IsAnagram(s string, t string) bool {
	sMap := make(map[rune]int, 0)
	tMap := make(map[rune]int, 0)

	if len(s) != len(t) {
		return false
	}

	for _, v := range s {
		sMap[v] += 1
	}

	for _, v := range t {
		tMap[v] += 1
	}

	for k := range sMap {
		if sMap[k] != tMap[k] {
			return false
		}
	}

	return true
}
