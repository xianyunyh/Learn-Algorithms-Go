package stack

// 利用栈 当是左括号时候，入栈，右括号出栈，进行匹配。
func isValid(s string) bool {
	stack := make([]string, 0)
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == "" {
			continue
		}
		if char == "[" || char == "{" || char == "(" {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			prev := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if char == "}" && prev != "{" {
				return false
			}
			if char == "]" && prev != "[" {
				return false
			}
			if char == ")" && prev != "(" {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
