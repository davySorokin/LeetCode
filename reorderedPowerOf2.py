class Solution:
    def reorderedPowerOf2(self, n: int) -> bool:
        # Generate all powers of 2 that are <= 10^9
        powers_of_2 = set()
        i = 1
        while i <= 10**9:
            powers_of_2.add(''.join(sorted(str(i))))
            i *= 2
        
        # Check if the sorted digits of n match any of the powers of 2
        return ''.join(sorted(str(n))) in powers_of_2
