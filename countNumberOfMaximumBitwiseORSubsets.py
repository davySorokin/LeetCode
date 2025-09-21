from typing import List

class Solution:
    def countMaxOrSubsets(self, nums: List[int]) -> int:
        n = len(nums)
        max_or = 0
        for v in nums:
            max_or |= v

        # tail_or[i] = OR of nums[i:]
        tail_or = [0]*(n+1)
        for i in range(n-1, -1, -1):
            tail_or[i] = tail_or[i+1] | nums[i]

        def dfs(i: int, cur: int) -> int:
            if i == n:
                return 1 if cur == max_or else 0
            # prune: even with all remaining, can't reach max_or
            if (cur | tail_or[i]) < max_or:
                return 0
            return dfs(i+1, cur) + dfs(i+1, cur | nums[i])

        return dfs(0, 0)

