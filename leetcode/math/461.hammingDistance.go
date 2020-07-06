package math
package math
func hammingDistance(x int, y int) int {
    return bitCount(x ^ y)
}

func bitCount(a int) int {
    num := 0
    for a != 0 {
        a = (a-1) & a
        num++
    }
    return num
}