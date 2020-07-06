package math

import "strconv"

func t(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		start = start * 10
		digit += 1
		count = 9 * start * digit
	}
	num := start + (n - 1)
	s := strconv.Itoa(num)
	return int(s[(num-1)%digit] - '0')
}
