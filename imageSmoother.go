func imageSmoother(img [][]int) [][]int {
    m, n := len(img), len(img[0])
    res := make([][]int, m)
    
    for i := range res {
        res[i] = make([]int, n)
        for j := range res[i] {
            sum, count := 0, 0
            for dx := -1; dx <= 1; dx++ {
                for dy := -1; dy <= 1; dy++ {
                    ni, nj := i + dx, j + dy
                    if ni >= 0 && ni < m && nj >= 0 && nj < n {
                        sum += img[ni][nj]
                        count++
                    }
                }
            }
            res[i][j] = sum / count
        }
    }
    
    return res
}
