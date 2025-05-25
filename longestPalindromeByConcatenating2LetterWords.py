from collections import Counter
from typing import List

class Solution:
    def longestPalindrome(self, words: List[str]) -> int:
        count = Counter(words)
        length = 0
        central_used = False

        for word in list(count.keys()):
            reversed_word = word[::-1]

            if word == reversed_word:
                # Pair up same-letter words like "gg"
                pairs = count[word] // 2
                length += pairs * 4
                count[word] -= pairs * 2
                if not central_used and count[word] > 0:
                    length += 2  # one central word
                    central_used = True
            else:
                if reversed_word in count:
                    pairs = min(count[word], count[reversed_word])
                    length += pairs * 4
                    count[word] -= pairs
                    count[reversed_word] -= pairs

        return length
