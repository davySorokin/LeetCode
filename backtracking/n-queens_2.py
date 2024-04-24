class Solution:
    def totalNQueens(self, n: int) -> int:
        results = 0  # Count of solutions found

        def backtrack(row=0):
            # Base case: if all rows are filled, increment results
            if row == n:
                nonlocal results
                results += 1
                return

            # Try placing queen in each column of current row
            for col in range(n):
                if is_safe(row, col):
                    cols[row] = col  # Mark column as occupied
                    backtrack(row + 1)  # Recur for next row
                    cols[row] = -1  # Backtrack: remove queen from column

        def is_safe(row, col):
            # Check for conflicts in columns and diagonals
            for i in range(row):
                if cols[i] == col or abs(cols[i] - col) == abs(i - row):
                    return False
            return True

        cols = [-1] * n  # Keeps track of queen placement in each column (-1 means empty)
        backtrack()
        return results
