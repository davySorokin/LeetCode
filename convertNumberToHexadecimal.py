class Solution:
    def toHex(self, num: int) -> str:
        if num == 0:
            return "0"
        
        hex_chars = "0123456789abcdef"
        result = ""

        # Handle two's complement for negative numbers
        if num < 0:
            num += 2 ** 32  # Convert to unsigned 32-bit integer

        while num > 0:
            digit = num % 16  # Get the last hex digit
            result = hex_chars[digit] + result  # Append it to the result
            num //= 16  # Move to the next digit

        return result
