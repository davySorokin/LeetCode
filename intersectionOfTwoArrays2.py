from typing import List
from collections import Counter

class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        # Count occurrences of elements in nums1 and nums2
        counts1 = Counter(nums1)
        counts2 = Counter(nums2)
        
        # Find the intersection by taking the minimum count of each element in both arrays
        intersection = []
        for num in counts1:
            if num in counts2:
                intersection.extend([num] * min(counts1[num], counts2[num]))
        
        return intersection
