from collections import Counter

class Solution:
    def firstUniqChar(self, s: str) -> int:
        # Count the frequency of each character
        char_count = Counter(s)
        
        # Find the first character with a count of 1
        for index, char in enumerate(s):
            if char_count[char] == 1:
                return index
        
        return -1
