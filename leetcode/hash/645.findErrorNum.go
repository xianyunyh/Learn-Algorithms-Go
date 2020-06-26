package hash

func findErrorNums(nums []int) []int {
    table := make(map[int]int)

    for _,v := range nums {
        if _,ok := table[v];ok {
            table[v] += 1
        }else {
            table[v] = 1
        }
    }
    repeat := 1
    mis := 1
    for i := 1; i<= len(nums);i++ {
        if _,ok := table[i];!ok {
            mis = i
        }
        if table[i] == 2 {
            repeat = i
        }
    } 
    return []int{repeat,mis}
}