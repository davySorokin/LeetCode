from typing import List

class Solution:
    def minFallingPathSum(self, grid: List[List[int]]) -> int:
        n = len(grid)
        if n == 1: return grid[0][0]  # Special case where there is only one element in the grid
        
        # Initialise the dp matrix
        dp = [[0] * n for _ in range(n)]
        dp[0] = grid[0][:]
        
        for i in range(1, n):
            # Find two smallest values from the previous row
            first_min, second_min = float('inf'), float('inf')
            for num in dp[i - 1]:
                if num < first_min:
                    first_min, second_min = num, first_min
                elif num < second_min:
                    second_min = num
            
            # Update the dp values for the current row
            for j in range(n):
                if dp[i - 1][j] != first_min:
                    dp[i][j] = first_min + grid[i][j]
                else:
                    dp[i][j] = second_min + grid[i][j]
        
        # answer is the minimum value in the last row of dp
        return min(dp[-1])
