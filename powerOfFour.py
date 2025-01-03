class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        # Check if n is greater than 0 and is a power of 2 (only one bit set in binary representation)
        # Also, ensure that the single bit set is at an odd position
        return n > 0 and (n & (n - 1)) == 0 and (n & 0xAAAAAAAA) == 0
