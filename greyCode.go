func grayCode(n int) []int {
    size := 1 << n
    res := make([]int, size)
    for i := 0; i < size; i++ {
        res[i] = i ^ (i >> 1)
    }
    return res
}
