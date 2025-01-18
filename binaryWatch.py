from typing import List

class Solution:
    def readBinaryWatch(self, turnedOn: int) -> List[str]:
        result = []
        
        # Iterate through all possible hour (0-11) and minute (0-59) combinations
        for hour in range(12):
            for minute in range(60):
                # Count the number of '1's in the binary representation
                if (bin(hour) + bin(minute)).count('1') == turnedOn:
                    # Format time correctly with leading zero for minutes
                    result.append(f"{hour}:{minute:02d}")
        
        return result
