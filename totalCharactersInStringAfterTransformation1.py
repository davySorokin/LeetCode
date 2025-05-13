class Solution:
    def lengthAfterTransformations(self, s: str, t: int) -> int:
        from collections import Counter, defaultdict

        MOD = 10**9 + 7

        # Count frequency of each character in the initial string
        freq = Counter(s)

        for _ in range(t):
            new_freq = defaultdict(int)
            for ch, count in freq.items():
                if ch == 'z':
                    new_freq['a'] += count
                    new_freq['b'] += count
                else:
                    next_ch = chr(ord(ch) + 1)
                    new_freq[next_ch] += count
            freq = new_freq

        return sum(freq.values()) % MOD
