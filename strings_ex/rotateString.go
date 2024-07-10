package stringsex

func RotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	buff := ""
	for i := 0; i < len(s); i++ {
		buff = s[i:] + s[:i]
		if buff == goal {
			return true
		}
	}

	return false
}
