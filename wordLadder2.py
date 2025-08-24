from collections import defaultdict, deque
from typing import List

class Solution:
    def findLadders(self, beginWord: str, endWord: str, wordList: List[str]) -> List[List[str]]:
        word_set = set(wordList)
        if endWord not in word_set:
            return []

        parents = defaultdict(list)  # child -> list of parents on shortest layers
        level = {beginWord}
        visited = set([beginWord])
        found = False

        while level and not found:
            next_level = set()
            for word in level:
                wchars = list(word)
                for i in range(len(wchars)):
                    orig = wchars[i]
                    for c in map(chr, range(ord('a'), ord('z') + 1)):
                        if c == orig: 
                            continue
                        wchars[i] = c
                        nxt = ''.join(wchars)
                        if nxt in word_set and nxt not in visited:
                            parents[nxt].append(word)
                            if nxt not in next_level:
                                next_level.add(nxt)
                            if nxt == endWord:
                                found = True
                    wchars[i] = orig
            visited |= next_level
            level = next_level

        if endWord not in parents:
            return []

        res, path = [], [endWord]

        def dfs(word: str):
            if word == beginWord:
                res.append(path[::-1])
                return
            for p in parents[word]:
                path.append(p)
                dfs(p)
                path.pop()

        dfs(endWord)
        return res
