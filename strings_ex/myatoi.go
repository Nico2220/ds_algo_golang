package stringsex

func MyAtoi(s string) int {
	newString := ""
	for _, v := range s {
		if string(v) == " " {
			continue
		}
		newString += string(v)
	}

	// sign := newString[0]

	i := 0
	for i < len(newString) {
		if string(newString[i]) != "0" {
			break
		}
		i++
	}
	newString = newString[i:]

	// j := len(newString) - 1
	// for  j  > 0 {
	// 	if string(newString[j]) != "0" {
	// 		break
	// 	}
	// 	j--
	// }
	// newString = newString[:j + 1]

	return 0
}
