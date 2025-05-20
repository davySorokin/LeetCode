from typing import List

class Solution:
    def isZeroArray(self, nums: List[int], queries: List[List[int]]) -> bool:
        n = len(nums)
        count = [0] * (n + 1)  # Difference array technique

        # Count how many times each index is covered
        for l, r in queries:
            count[l] += 1
            if r + 1 < len(count):
                count[r + 1] -= 1

        # Compute prefix sum to get actual coverage count at each index
        for i in range(1, n):
            count[i] += count[i - 1]

        # Compare count with nums
        for i in range(n):
            if nums[i] > count[i]:
                return False

        return True
