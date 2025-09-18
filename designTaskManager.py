from typing import List, Dict, Tuple
import heapq

class TaskManager:

    def __init__(self, tasks: List[List[int]]):
        # taskId -> (userId, priority) for currently active tasks
        self.info: Dict[int, Tuple[int, int]] = {}
        # max-heap simulated via negatives; store ( -priority, -taskId, taskId )
        self.heap = []
        for userId, taskId, priority in tasks:
            self.info[taskId] = (userId, priority)
            heapq.heappush(self.heap, (-priority, -taskId, taskId))

    def add(self, userId: int, taskId: int, priority: int) -> None:
        self.info[taskId] = (userId, priority)
        heapq.heappush(self.heap, (-priority, -taskId, taskId))

    def edit(self, taskId: int, newPriority: int) -> None:
        userId, _ = self.info[taskId]
        self.info[taskId] = (userId, newPriority)
        heapq.heappush(self.heap, (-newPriority, -taskId, taskId))

    def rmv(self, taskId: int) -> None:
        # Lazy delete: just remove from info; stale heap entries will be skipped
        if taskId in self.info:
            del self.info[taskId]

    def execTop(self) -> int:
        # Pop until we find a live entry that matches current priority
        while self.heap:
            nprio, ntid, tid = heapq.heappop(self.heap)
            if tid not in self.info:
                continue  # removed
            userId, cur_prio = self.info[tid]
            if -nprio == cur_prio:
                # valid top; execute and remove
                del self.info[tid]
                return userId
            # else: stale due to edit; skip
        return -1
