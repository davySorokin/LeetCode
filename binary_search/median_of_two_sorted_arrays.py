from typing import List

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        # Ensure nums1 is the shorter array
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        
        m, n = len(nums1), len(nums2)
        total = m + n
        half = (total + 1) // 2
        
        # Binary search on the shorter array
        left, right = 0, m
        while left <= right:
            i = (left + right) // 2
            j = half - i
            
            nums1LeftMax = float('-inf') if i == 0 else nums1[i - 1]
            nums1RightMin = float('inf') if i == m else nums1[i]
            nums2LeftMax = float('-inf') if j == 0 else nums2[j - 1]
            nums2RightMin = float('inf') if j == n else nums2[j]
            
            if nums1LeftMax <= nums2RightMin and nums2LeftMax <= nums1RightMin:
                if total % 2 == 1:
                    return max(nums1LeftMax, nums2LeftMax)
                return (max(nums1LeftMax, nums2LeftMax) + min(nums1RightMin, nums2RightMin)) / 2
            elif nums1LeftMax > nums2RightMin:
                right = i - 1
            else:
                left = i + 1
        
        raise ValueError("The input arrays are not sorted or not following the problem constraints.")
