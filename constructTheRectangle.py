from typing import List
import math

class Solution:
    def constructRectangle(self, area: int) -> List[int]:
        # Start from the square root of the area and move down to find the largest factor
        w = int(math.sqrt(area))
        while area % w != 0:
            w -= 1
        l = area // w
        return [l, w]

sol = Solution()
print(sol.constructRectangle(4))       # Output: [2, 2]
print(sol.constructRectangle(37))      # Output: [37, 1]
print(sol.constructRectangle(122122))  # Output: [427, 286]
