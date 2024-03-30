from typing import List

class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:

        # These arrays are used to get row and column numbers of 8 neighbors
        # of a given cell
        direction_x = [-1, -1, -1, 0, 0, 1, 1, 1]
        direction_y = [-1, 0, 1, -1, 1, -1, 0, 1]

        m, n = len(board), len(board[0])
        
        # Create a copy of the original board
        copy_board = [[board[row][col] for col in range(n)] for row in range(m)]

        # Iterate through board cell by cell.
        for row in range(m):
            for col in range(n):

                # For each cell count the number of live neighbors.
                live_neighbors = 0
                for k in range(8):
                    new_row, new_col = row + direction_x[k], col + direction_y[k]
                    if (new_row < m and new_row >= 0) and (new_col < n and new_col >= 0) and (copy_board[new_row][new_col] == 1):
                        live_neighbors += 1

                # Rule 1 or Rule 3
                if copy_board[row][col] == 1 and (live_neighbors < 2 or live_neighbors > 3):
                    board[row][col] = 0
                # Rule 4
                if copy_board[row][col] == 0 and live_neighbors == 3:
                    board[row][col] = 1
                # Rule 2 is implicitly handled as we don't change the cell when it has 2 or 3 neighbors.

# Example usage:
# solution = Solution()
# example_board = [
#    [0, 1, 0],
#    [0, 0, 1],
#    [1, 1, 1],
#    [0, 0, 0]
# ]
# solution.gameOfLife(example_board)
# print(example_board) # Expected output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
