from typing import List
from collections import defaultdict

class Solution:
    def sortMatrix(self, grid: List[List[int]]) -> List[List[int]]:
        n = len(grid)
        diags = defaultdict(list)

        # Collect values by diagonal (i - j)
        for i in range(n):
            for j in range(n):
                diags[i - j].append(grid[i][j])

        # Sort diagonals depending on region
        for d in diags:
            if d >= 0:  # bottom-left incl main diag
                diags[d].sort(reverse=True)  # non-increasing
            else:       # top-right
                diags[d].sort()              # non-decreasing

        # Write back to grid
        for i in range(n):
            for j in range(n):
                grid[i][j] = diags[i - j].pop(0)

        return grid
