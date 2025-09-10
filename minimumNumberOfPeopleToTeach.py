from typing import List
from collections import defaultdict

class Solution:
    def minimumTeachings(self, n: int, languages: List[List[int]], friendships: List[List[int]]) -> int:
        m = len(languages)
        # Convert user language lists to sets for fast lookup
        user_languages = [set(langs) for langs in languages]
        
        # Identify all friend pairs that cannot communicate
        to_fix = []
        for u, v in friendships:
            u -= 1
            v -= 1
            if user_languages[u].intersection(user_languages[v]):
                continue
            to_fix.append((u, v))
        
        # For each language, count how many unique users among those in to_fix need to learn it
        min_teach = float('inf')
        for lang in range(1, n + 1):
            teach_set = set()
            for u, v in to_fix:
                if lang not in user_languages[u]:
                    teach_set.add(u)
                if lang not in user_languages[v]:
                    teach_set.add(v)
            min_teach = min(min_teach, len(teach_set))
        
        return 0 if not to_fix else min_teach
