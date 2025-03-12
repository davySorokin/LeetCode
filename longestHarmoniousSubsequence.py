from typing import List
from collections import Counter

class Solution:
    def findLHS(self, nums: List[int]) -> int:
        count = Counter(nums)
        max_length = 0

        for num in count:
            if num + 1 in count:  # Check if there is a number that is exactly 1 greater
                max_length = max(max_length, count[num] + count[num + 1])

        return max_length
