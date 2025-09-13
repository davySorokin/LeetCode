from collections import Counter

class Solution:
    def maxFreqSum(self, s: str) -> int:
        vowels = {'a', 'e', 'i', 'o', 'u'}
        freq = Counter(s)
        
        max_vowel = max((freq[char] for char in vowels if char in freq), default=0)
        max_consonant = max((freq[char] for char in freq if char not in vowels), default=0)
        
        return max_vowel + max_consonant
