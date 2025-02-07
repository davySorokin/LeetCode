from typing import List

class Solution:
    def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
        if not timeSeries:
            return 0
        
        total_duration = 0

        for i in range(len(timeSeries) - 1):
            total_duration += min(timeSeries[i + 1] - timeSeries[i], duration)
        
        return total_duration + duration
