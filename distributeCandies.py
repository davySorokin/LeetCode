from typing import List

class Solution:
    def distributeCandies(self, candyType: List[int]) -> int:
        unique_candies = len(set(candyType))  # Count unique candy types
        max_allowed = len(candyType) // 2  # Maximum candies Alice can eat
        return min(unique_candies, max_allowed)  # At most min of both
