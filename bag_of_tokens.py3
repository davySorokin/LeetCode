from typing import List

class Solution:
    def bagOfTokensScore(self, tokens: List[int], power: int) -> int:
        tokens.sort()
        score = 0
        max_score = 0
        left, right = 0, len(tokens) - 1

        while left <= right:
            if power >= tokens[left]:  # Play the smallest token face-up.
                power -= tokens[left]
                score += 1
                left += 1
                max_score = max(max_score, score)
            elif score > 0:  # Play the largest token face-down.
                power += tokens[right]
                score -= 1
                right -= 1
            else:  # Can't play any token.
                break
        
        return max_score
