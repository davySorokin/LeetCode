from typing import List

class Solution:
    def countBits(self, n: int) -> List[int]:
        # Initialise the result array with 0
        ans = [0] * (n + 1)
        
        # Use dynamic programming; the number of 1s in `i` is:
        # The number of 1s in `i >> 1` + (i % 2)
        for i in range(1, n + 1):
            ans[i] = ans[i >> 1] + (i & 1)
        
        return ans
