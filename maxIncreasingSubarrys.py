from typing import List

class Solution:
    def maxIncreasingSubarrays(self, nums: List[int]) -> int:
        n = len(nums)
        left = [1] * n
        right = [1] * n

        # compute increasing lengths ending at i
        for i in range(1, n):
            if nums[i] > nums[i - 1]:
                left[i] = left[i - 1] + 1

        # compute increasing lengths starting at i
        for i in range(n - 2, -1, -1):
            if nums[i] < nums[i + 1]:
                right[i] = right[i + 1] + 1

        ans = 0
        # check adjacent subarrays (split between i and i+1)
        for i in range(n - 1):
            ans = max(ans, min(left[i], right[i + 1]))

        return ans


# tests
# print(Solution().maxIncreasingSubarrays([2,5,7,8,9,2,3,4,3,1]))  # 3
# print(Solution().maxIncreasingSubarrays([1,2,3,4,4,4,4,5,6,7]))  # 2
