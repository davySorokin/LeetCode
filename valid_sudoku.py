class Solution:
    def isValidSudoku(self, board):
        # Use hash set to record the status
        rows = [set() for _ in range(9)]
        cols = [set() for _ in range(9)]
        boxes = [set() for _ in range(9)]

        for i in range(9):
            for j in range(9):
                # Check if the position is filled with a number
                if board[i][j] != '.':
                    num = board[i][j]
                    box_index = (i // 3) * 3 + j // 3
                    
                    # Check if the number is already in the current row, column, or box
                    if num in rows[i] or num in cols[j] or num in boxes[box_index]:
                        return False
                    
                    # Add the number to the current row, column, and box
                    rows[i].add(num)
                    cols[j].add(num)
                    boxes[box_index].add(num)
                    
        return True
