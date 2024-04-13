class Solution:
    def solve(self, board):
        if not board or not board[0]:
            return
        
        def dfs(i, j):
            if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]) or board[i][j] != 'O':
                return
            board[i][j] = 'T'
            dfs(i+1, j)
            dfs(i-1, j)
            dfs(i, j+1)
            dfs(i, j-1)
        
        # Run DFS on 'O's on the border and mark them as 'T' (temporarily safe)
        for i in range(len(board)):
            for j in range(len(board[0])):
                if (i in [0, len(board)-1] or j in [0, len(board[0])-1]) and board[i][j] == 'O':
                    dfs(i, j)
        
        # Flip the remaining 'O's to 'X' and 'T's back to 'O'
        for i in range(len(board)):
            for j in range(len(board[0])):
                if board[i][j] == 'O':
                    board[i][j] = 'X'
                elif board[i][j] == 'T':
                    board[i][j] = 'O'
