package lcoffer

func translateNum(num int) int {
	a, b := 1, 1
	x, y := num%10, num%10 //记录当前位和后一位
	for num != 0 {
		num = num / 10
		x = num % 10
		temp := 10*x + y
		c := a
		if temp >= 10 && temp <= 25 {
			c = a + b
		}
		b = a
		a = c
		y = x
	}
	return a
}
