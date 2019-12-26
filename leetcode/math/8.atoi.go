package math

func myAtoi(str string) int {
	max := int(^uint32(0) >> 1)
	min := ^max
	isNum := func(c uint8) bool {
		return '0' <= c && c <= '9'
	}
	isFlag := func(c uint8) bool {
		return c == '+' || c == '-'
	}
	// 去除前后空格
	i := 0
	for ; i < len(str); i++ {
		if str[i] != ' ' {
			break
		}
	}
	//截断字符
	str = str[i:]
	l := len(str)
	if l == 0 {
		return 0
	}
	if !isFlag(str[0]) && !isNum(str[0]) {
		return 0
	}
	f := 1
	if isFlag(str[0]) {
		if str[0] == '-' {
			f = -1
		}
		str = str[1:]
		l = len(str)
	}
	res := 0
	for i := 0; i < l; i++ {
		if !isNum(str[i]) {
			break
		}
		res += int(str[i] - '0')
		if f*res > max {
			return max
		}
		if f*res < min {
			return min
		}
		res *= 10
	}
	res = f * res / 10
	return res
}
