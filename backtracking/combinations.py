from typing import List

class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        results = []
        
        def backtrack(start, combination):
            # If the combination is complete
            if len(combination) == k:
                results.append(combination[:])
                return
            
            # Go through all numbers from start to n
            for i in range(start, n + 1):
                # Include i into the current combination
                combination.append(i)
                # Move on to the next element
                backtrack(i + 1, combination)
                # Exclude i from the combination (backtrack)
                combination.pop()
        
        backtrack(1, [])
        return results
