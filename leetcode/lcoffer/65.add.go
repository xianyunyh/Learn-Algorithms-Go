package lcoffer

// 无进位和 与 异或运算 规律相同， a ^ b
// 进位 和 与运算 规律相同（并需左移一位） (a&b) << 1
func add(a int, b int) int {
    for b != 0 {
        c := (a & b) << 1
        a ^= b
        b = c
    }
    return a
}