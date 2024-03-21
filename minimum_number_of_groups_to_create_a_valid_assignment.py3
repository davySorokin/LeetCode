class Solution:
    def minGroupsForValidAssignment(self, balls: List[int]) -> int:
        # Calculate the frequency of each number.
        freq = Counter(balls)
        # The maximum frequency will determine the range we need to iterate over.
        max_freq = max(freq.values())
        
        # This will hold the minimum number of groups found.
        min_groups = float('inf')
        
        # Iterate over each possible group size from 1 to the maximum frequency.
        for x in range(1, max_freq + 1):
            groups = 0
            for f in freq.values():
                # Calculate the number of groups of size x and x + 1 needed.
                a = f // (x + 1)
                b = f % (x + 1)
                if b == 0:
                    # We can form groups all of size x+1.
                    groups += a
                elif x - b <= a:
                    # We can form groups of size x and x+1.
                    groups += a + 1
                else:
                    # It's impossible to form groups with this size, skip to next.
                    groups = float('inf')
                    break
            # Update the minimum groups if this is the lowest count we've found.
            min_groups = min(min_groups, groups)
        
        return min_groups
