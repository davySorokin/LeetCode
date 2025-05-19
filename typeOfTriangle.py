from typing import List

class Solution:
    def triangleType(self, nums: List[int]) -> str:
        a, b, c = sorted(nums)
        # Check for triangle inequality
        if a + b <= c:
            return "none"
        # Check for equilateral triangle
        if a == b == c:
            return "equilateral"
        # Check for isosceles triangle
        if a == b or b == c or a == c:
            return "isosceles"
        # If all sides are different
        return "scalene"
