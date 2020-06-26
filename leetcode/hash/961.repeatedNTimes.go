func repeatedNTimes(A []int) int {
    table := make(map[int]int)

    for _,v := range A {
        if _,ok :=  table[v];ok {
            table[v] = table[v] +1
        }else {
            table[v] = 1
        }
    }
    for k,v  := range table {
        if v > 1 {
            return k
        }
    }
    return -1
}