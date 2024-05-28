func solveSudoku(board [][]byte) {
    if solve(board) {
        return
    }
}

func solve(board [][]byte) bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                for c := '1'; c <= '9'; c++ {
                    if isValid(board, i, j, byte(c)) {
                        board[i][j] = byte(c)
                        if solve(board) {
                            return true
                        }
                        board[i][j] = '.'
                    }
                }
                return false
            }
        }
    }
    return true
}

func isValid(board [][]byte, row, col int, c byte) bool {
    for i := 0; i < 9; i++ {
        // Check row
        if board[i][col] == c {
            return false
        }
        // Check column
        if board[row][i] == c {
            return false
        }
        // Check 3x3 box
        boxRow := 3 * (row / 3) + i / 3
        boxCol := 3 * (col / 3) + i % 3
        if board[boxRow][boxCol] == c {
            return false
        }
    }
    return true
}
