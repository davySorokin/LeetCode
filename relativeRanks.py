from typing import List

class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:

        sorted_scores = sorted(enumerate(score), key=lambda x: x[1], reverse=True)
        
        result = [""] * len(score)

        for rank, (idx, _) in enumerate(sorted_scores):
            if rank == 0:
                result[idx] = "Gold Medal"
            elif rank == 1:
                result[idx] = "Silver Medal"
            elif rank == 2:
                result[idx] = "Bronze Medal"
            else:
                result[idx] = str(rank + 1)

        return result
