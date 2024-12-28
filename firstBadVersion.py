# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        left, right = 1, n
        while left < right:
            mid = left + (right - left) // 2  # To prevent overflow
            if isBadVersion(mid):
                right = mid  # Keep checking in the left part
            else:
                left = mid + 1  # Check the right part
        return left  # At this point, left == right and is the first bad version
