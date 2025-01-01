class Solution:
    def isPowerOfThree(self, n: int) -> bool:
        # The maximum power of 3 value that fits in a 32-bit signed integer is 3^19 = 1162261467
        # Any number that is a power of 3 must divide this number without leaving a remainder
        return n > 0 and 1162261467 % n == 0
