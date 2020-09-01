package string

//利用双指针
//回文串有可能是奇数、偶数 需要两次查找，从中心向两边发散
// abba abcba
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		temp := Palindrome(s, i, i)
		if len(temp) > len(res) {
			res = temp
		}
		temp = Palindrome(s, i, i+1)
		if len(temp) > len(res) {
			res = temp
		}
	}
	return res
}

func Palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
