class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        charIndexMap = {}
        maxLength = 0
        start = 0
        
        for i, char in enumerate(s):
            # If the character is in the map and its index is within the current window
            if char in charIndexMap and charIndexMap[char] >= start:
                start = charIndexMap[char] + 1  # Move the start of the window
            charIndexMap[char] = i
            maxLength = max(maxLength, i - start + 1)  # Update the max length
        
        return maxLength
