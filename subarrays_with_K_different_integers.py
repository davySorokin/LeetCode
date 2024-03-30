from typing import List

class Solution:
    def subarraysWithKDistinct(self, nums: List[int], k: int) -> int:
        def atMostKDistinct(k):
            count = {}
            res = i = 0
            for j in range(len(nums)):
                if nums[j] not in count:
                    k -= 1
                count[nums[j]] = count.get(nums[j], 0) + 1
                while k < 0:
                    count[nums[i]] -= 1
                    if count[nums[i]] == 0:
                        del count[nums[i]]
                        k += 1
                    i += 1
                res += j - i + 1
            return res

        # To find the exact K distinct, we can find at most K distinct - at most (K-1) distinct.
        return atMostKDistinct(k) - atMostKDistinct(k - 1)

# Example:
# sol = Solution()
# example_nums = [1, 2, 1, 2, 3]
# example_k = 2
# print(sol.subarraysWithKDistinct(example_nums, example_k))  # Expected output is 7
