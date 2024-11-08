func findColumnWidth(grid [][]int) []int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return []int{}
    }
    
    m, n := len(grid), len(grid[0])
    result := make([]int, n)
    
    for col := 0; col < n; col++ {
        maxLen := 0
        for row := 0; row < m; row++ {
            numLen := len(fmt.Sprintf("%d", grid[row][col]))
            if numLen > maxLen {
                maxLen = numLen
            }
        }
        result[col] = maxLen
    }
    
    return result
}
