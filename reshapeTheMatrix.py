from typing import List

class Solution:
    def matrixReshape(self, mat: List[List[int]], r: int, c: int) -> List[List[int]]:
        # Get the number of rows and columns in the original matrix
        m, n = len(mat), len(mat[0])
        
        if m * n != r * c:
            return mat  # Return original matrix if reshape isn't possible
        
        flat = [num for row in mat for num in row]
        
        return [flat[i * c:(i + 1) * c] for i in range(r)]
