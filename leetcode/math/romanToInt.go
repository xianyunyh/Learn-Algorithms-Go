package math

// [13] 罗马数字转整数
func romanToInt(s string) int {
	nums := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	lens := len(s)
	num := 0

	for i := 0; i < lens; i++ {
		c := string(s[i])
		if i == lens-1 {
			num += nums[c]
			continue
		}
		if c == "I" || c == "X" || c == "C" {
			c1 := c + string(s[i+1])
			if v, ok := nums[c1]; ok {
				num += v
				i = i + 1
				continue
			}
		}
		num += nums[c]
	}
	return num
}
