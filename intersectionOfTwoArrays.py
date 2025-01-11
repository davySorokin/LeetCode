from typing import List

class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        # Use set intersection to get unique elements in both arrays
        return list(set(nums1) & set(nums2))
