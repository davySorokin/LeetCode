func areSimilar(mat [][]int, k int) bool {
    m := len(mat)
    n := len(mat[0])

    // Helper function to shift a row to the left
    leftShift := func(row []int, steps int) []int {
        steps = steps % n
        return append(row[steps:], row[:steps]...)
    }

    // Helper function to shift a row to the right
    rightShift := func(row []int, steps int) []int {
        steps = steps % n
        return append(row[n-steps:], row[:n-steps]...)
    }

    // Create a copy of the original matrix
    original := make([][]int, m)
    for i := range mat {
        original[i] = append([]int(nil), mat[i]...)
    }

    // Perform k steps of shifts
    for step := 0; step < k; step++ {
        for i := 0; i < m; i++ {
            if i%2 == 0 {
                mat[i] = leftShift(mat[i], 1)
            } else {
                mat[i] = rightShift(mat[i], 1)
            }
        }
    }

    // Check if the final matrix is the same as the original
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] != original[i][j] {
                return false
            }
        }
    }

    return true
}
