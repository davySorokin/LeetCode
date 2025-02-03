class Solution:
    def findComplement(self, num: int) -> int:
        # Get the bit length of the number
        bit_length = num.bit_length()
        
        # Create a mask with all bits set to 1 of the same length as num
        mask = (1 << bit_length) - 1
        
        # XOR num with mask to get the complement
        return num ^ mask

solution = Solution()
print(solution.findComplement(5))  # Output: 2
print(solution.findComplement(1))  # Output: 0
