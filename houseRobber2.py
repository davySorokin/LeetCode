from typing import List

class Solution:
    def rob(self, nums: List[int]) -> int:
        def rob_linear(houses: List[int]) -> int:
            prev = curr = 0
            for money in houses:
                prev, curr = curr, max(curr, prev + money)
            return curr

        n = len(nums)
        if n == 1:
            return nums[0]
        
        return max(rob_linear(nums[:-1]), rob_linear(nums[1:]))
