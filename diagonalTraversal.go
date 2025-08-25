func findDiagonalOrder(mat [][]int) []int {
    m := len(mat)
    if m == 0 {
        return []int{}
    }
    n := len(mat[0])
    var result []int

    for d := 0; d < m+n-1; d++ {
        var intermediate []int
        var r, c int

        if d < n {
            r = 0
            c = d
        } else {
            r = d - n + 1
            c = n - 1
        }

        for r < m && c >= 0 {
            intermediate = append(intermediate, mat[r][c])
            r++
            c--
        }

        if d%2 == 0 {
            // Reverse the diagonal
            for i := len(intermediate) - 1; i >= 0; i-- {
                result = append(result, intermediate[i])
            }
        } else {
            result = append(result, intermediate...)
        }
    }

    return result
}
