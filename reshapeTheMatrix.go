func matrixReshape(mat [][]int, r int, c int) [][]int {
    m, n := len(mat), len(mat[0])

    // Check if reshaping is possible
    if m*n != r*c {
        return mat
    }

    // Create a flattened list of elements
    flattened := make([]int, 0, m*n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            flattened = append(flattened, mat[i][j])
        }
    }

    // Construct the reshaped matrix
    reshaped := make([][]int, r)
    for i := 0; i < r; i++ {
        reshaped[i] = make([]int, c)
        for j := 0; j < c; j++ {
            reshaped[i][j] = flattened[i*c + j]
        }
    }
    return reshaped
}
