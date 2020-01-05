package hash

// 我们可以利用一个 map 来处理映射。对于 s 到 t 的映射，我们同时遍历 s 和 t ，假设当前遇到的字母分别是 c1 和 c2 。

// 如果 map[c1] 不存在，那么就将 c1 映射到 c2 ，即 map[c1] = c2。

// 如果 map[c1] 存在，那么就判断 map[c1] 是否等于 c2，也就是验证之前的映射和当前的字母是否相同。
//
func isIsomorphic(s string, t string) bool {
	check := func(s string, t string) bool {
		m := make(map[rune]rune)
		for i, v := range s {
			if _, ok := m[v]; ok {
				if m[v] != rune(t[i]) {
					return false
				}
			} else {
				m[v] = rune(t[i])
			}

		}
		return true
	}
	return check(s, t) && check(t, s)
}
