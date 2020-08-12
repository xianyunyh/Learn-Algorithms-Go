package lcoffer


func maxValue(grid [][]int) int {
    row,column := len(grid),len(grid[0])
    for j := 1;j < column;j++ {
        grid[0][j] += grid[0][j-1]
    }
    for i := 1;i < row;i++ {
        grid[i][0] += grid[i-1][0]
    }
    for i := 1;i < row;i++ {

        for j := 1;j < column;j++ {
            grid[i][j] +=  max(grid[i][j-1],grid[i-1][j])
        }
    }
    return grid[row-1][column-1]
}

func  max (a,b int) int {
    if a > b {
        return a
    }
    return b
}