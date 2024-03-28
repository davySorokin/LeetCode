from collections import Counter

class Solution:
    def findSubstring(self, s, words):
        if not s or not words:
            return []
        
        word_len = len(words[0])
        num_words = len(words)
        total_len = word_len * num_words
        word_count = Counter(words)
        
        def check(i):
            seen = Counter()
            for j in range(i, i + total_len, word_len):
                curr_word = s[j:j + word_len]
                if curr_word in word_count:
                    seen[curr_word] += 1
                    if seen[curr_word] > word_count[curr_word]:
                        return False
                else:
                    return False
            return True
        
        return [i for i in range(len(s) - total_len + 1) if check(i)]
