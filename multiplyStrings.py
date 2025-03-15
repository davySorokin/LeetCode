class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        
        # Result array to store multiplication results
        result = [0] * (len(num1) + len(num2))
        
        # Reverse both numbers to simulate manual multiplication
        num1 = num1[::-1]
        num2 = num2[::-1]
        
        # Multiply digit by digit
        for i in range(len(num1)):
            for j in range(len(num2)):
                digit_mul = (ord(num1[i]) - ord('0')) * (ord(num2[j]) - ord('0'))
                result[i + j] += digit_mul
                result[i + j + 1] += result[i + j] // 10  # Carry
                result[i + j] %= 10  # Remainder
        
        # Remove leading zeros
        while len(result) > 1 and result[-1] == 0:
            result.pop()
        
        # Convert result array to string
        return ''.join(str(digit) for digit in result[::-1])

# Example:
sol = Solution()
print(sol.multiply("123", "456"))
