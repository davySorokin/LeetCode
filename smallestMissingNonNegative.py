from collections import Counter
from typing import List

class Solution:
    def findSmallestInteger(self, nums: List[int], value: int) -> int:
        freq = Counter(((num % value) + value) % value for num in nums)
        mex = 0
        while True:
            rem = mex % value
            if freq[rem] > 0:
                freq[rem] -= 1
                mex += 1
            else:
                return mex
