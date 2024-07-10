package stringsex

func IsValid(s string) bool {
	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {

			if len(stack) == 0 {
				return false
			}

			value := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if v == ')' && value != '(' {
				return false
			} else if v == ']' && value != '[' {
				return false
			} else if v == '}' && value != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}
