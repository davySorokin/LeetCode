class Solution {
    func solveNQueens(_ n: Int) -> [[String]] {
        var result = [[String]]()
        var board = [String](repeating: String(repeating: ".", count: n), count: n)
        var columns = Set<Int>()
        var diag1 = Set<Int>()
        var diag2 = Set<Int>()
        
        func backtrack(row: Int) {
            if row == n {
                result.append(board)
                return
            }
            
            for col in 0..<n {
                if columns.contains(col) || diag1.contains(row - col) || diag2.contains(row + col) {
                    continue // Skip if there's a conflict
                }
                
                // Place the queen
                board[row] = String(repeating: ".", count: col) + "Q" + String(repeating: ".", count: n - col - 1)
                columns.insert(col)
                diag1.insert(row - col)
                diag2.insert(row + col)
                
                backtrack(row: row + 1) // Move to the next row
                
                // Backtrack: undo the placement
                columns.remove(col)
                diag1.remove(row - col)
                diag2.remove(row + col)
            }
        }
        
        backtrack(row: 0)
        return result
    }
}
