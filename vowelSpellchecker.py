from typing import List

class Solution:
    def spellchecker(self, wordlist: List[str], queries: List[str]) -> List[str]:
        vowels = set('aeiou')

        def devowel(word: str) -> str:
            return ''.join('*' if ch in vowels else ch for ch in word.lower())

        exact = set(wordlist)
        cap_map = {}
        vow_map = {}

        for word in wordlist:
            low = word.lower()
            vow = devowel(word)
            if low not in cap_map:
                cap_map[low] = word
            if vow not in vow_map:
                vow_map[vow] = word

        result = []
        for query in queries:
            if query in exact:
                result.append(query)
            elif query.lower() in cap_map:
                result.append(cap_map[query.lower()])
            elif devowel(query) in vow_map:
                result.append(vow_map[devowel(query)])
            else:
                result.append("")
        return result
