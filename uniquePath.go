func uniquePaths(m int, n int) int {
    a := m + n - 2
    b := min(m-1, n-1)
    res := 1
    for i := 1; i <= b; i++ {
        res = res * (a - b + i) / i
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
