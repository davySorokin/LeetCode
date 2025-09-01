from typing import List
import heapq

class Solution:
    def maxAverageRatio(self, classes: List[List[int]], extraStudents: int) -> float:
        # Gain if we add one student to (p, t)
        def gain(p: int, t: int) -> float:
            return (p + 1) / (t + 1) - p / t

        # Max-heap via negative keys: (-gain, p, t)
        heap = [(-gain(p, t), p, t) for p, t in classes]
        heapq.heapify(heap)

        for _ in range(extraStudents):
            g, p, t = heapq.heappop(heap)
            p += 1
            t += 1
            heapq.heappush(heap, (-gain(p, t), p, t))

        # Compute final average
        total = 0.0
        for _, p, t in heap:
            total += p / t
        return total / len(classes)
