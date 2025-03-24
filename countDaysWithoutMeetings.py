from typing import List

class Solution:
    def countDays(self, days: int, meetings: List[List[int]]) -> int:
        # Sort meetings by start time
        meetings.sort()
        
        # Merge overlapping intervals
        merged = []
        for start, end in meetings:
            if not merged or start > merged[-1][1]:
                merged.append([start, end])
            else:
                merged[-1][1] = max(merged[-1][1], end)
        
        # Calculate total days with meetings
        meeting_days = 0
        for start, end in merged:
            meeting_days += (end - start + 1)
        
        # Subtract from total days
        return days - meeting_days
