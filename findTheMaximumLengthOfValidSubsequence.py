from typing import List

class Solution:
    def maximumLength(self, nums: List[int]) -> int:
        even, odd = 0, 0  # Two counters for alternating sequences

        for num in nums:
            if num % 2 == 0:
                even = odd + 1
            else:
                odd = even + 1

        return max(even, odd)
