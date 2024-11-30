class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        result = []
        while columnNumber > 0:
            columnNumber -= 1  # To make it 0-indexed
            result.append(chr(columnNumber % 26 + ord('A')))
            columnNumber //= 26
        return ''.join(result[::-1])  # Reverse the result to get the correct title
