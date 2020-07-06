package math

import "strconv"

func findNthDigit(n int) int {
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start = start * 10
		digit += 1
		count = 9 * start * digit
	}
	num := start + (n-1)/digit
	s := strconv.Itoa(num)
	return int(s[(n-1)%digit] - '0')
}
