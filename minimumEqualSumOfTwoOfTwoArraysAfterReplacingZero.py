from typing import List

class Solution:
    def minSum(self, nums1: List[int], nums2: List[int]) -> int:
        sum1, sum2 = sum(nums1), sum(nums2)
        zero1, zero2 = nums1.count(0), nums2.count(0)
        min1 = sum1 + zero1
        min2 = sum2 + zero2
        if min1 == min2:
            return min1
        diff = abs(min1 - min2)
        zeros_to_increase = zero1 if min1 < min2 else zero2
        if zeros_to_increase == 0:
            return -1
        return max(min1, min2)
