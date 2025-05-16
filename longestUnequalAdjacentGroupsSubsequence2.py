from typing import List

class Solution:
    def getWordsInLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        n = len(words)
        dp = [1] * n
        prev = [-1] * n

        def hamming_distance(w1, w2):
            return sum(a != b for a, b in zip(w1, w2)) if len(w1) == len(w2) else float('inf')

        for i in range(n):
            for j in range(i):
                if groups[i] != groups[j] and len(words[i]) == len(words[j]) and hamming_distance(words[i], words[j]) == 1:
                    if dp[j] + 1 > dp[i]:
                        dp[i] = dp[j] + 1
                        prev[i] = j

        # Locate the index of max length subsequence
        max_len = max(dp)
        index = dp.index(max_len)

        # Then Reconstruct the sequence
        res = []
        while index != -1:
            res.append(words[index])
            index = prev[index]

        return res[::-1]
