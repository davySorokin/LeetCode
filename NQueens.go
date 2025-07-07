package main

import "fmt"

func solveNQueens(n int) [][]string {
    result := [][]string{}
    board := make([]string, n)
    columns := make([]bool, n)        // Columns where queens are placed
    diag1 := make([]bool, 2*n-1)      // Diagonals (\ direction)
    diag2 := make([]bool, 2*n-1)      // Diagonals (/ direction)
    
    var backtrack func(row int)
    backtrack = func(row int) {
        if row == n {
            solution := make([]string, n)
            copy(solution, board)
            result = append(result, solution)
            return
        }
        
        for col := 0; col < n; col++ {
            if columns[col] || diag1[row-col+n-1] || diag2[row+col] {
                continue // Skip if the queen can be attacked
            }
            
            // Place the queen
            board[row] = createRow(n, col)
            columns[col] = true
            diag1[row-col+n-1] = true
            diag2[row+col] = true
            
            // Move to next row
            backtrack(row + 1)
            
            // Backtrack, undo the placement
            columns[col] = false
            diag1[row-col+n-1] = false
            diag2[row+col] = false
        }
    }
    
    backtrack(0)
    return result
}

func createRow(n, col int) string {
    row := make([]rune, n)
    for i := 0; i < n; i++ {
        if i == col {
            row[i] = 'Q'
        } else {
            row[i] = '.'
        }
    }
    return string(row)
}

/*
func main() {
    n := 4
    solutions := solveNQueens(n)
    fmt.Println(solutions)
}
*/
