from collections import deque
from bisect import bisect_left, bisect_right
from typing import List

class Router:
    def __init__(self, memoryLimit: int):
        self.cap = memoryLimit
        self.q = deque()  # FIFO of (s, d, t)
        self.present = set()  # set of (s, d, t) for duplicate detection
        # destination -> {"ts": [timestamps...], "start": idx_of_first_valid}
        self.per_dest = {}

    def _ensure_dest(self, d: int):
        if d not in self.per_dest:
            self.per_dest[d] = {"ts": [], "start": 0}

    def _advance_dest_front(self, d: int):
        """Remove one packet (the oldest overall) from destination d."""
        b = self.per_dest[d]
        b["start"] += 1
        # occasional compaction to keep arrays small
        if b["start"] > 64 and b["start"] * 2 >= len(b["ts"]):
            b["ts"] = b["ts"][b["start"]:]
            b["start"] = 0

    def addPacket(self, source: int, destination: int, timestamp: int) -> bool:
        key = (source, destination, timestamp)
        if key in self.present:
            return False

        # Evict oldest if at capacity
        if len(self.q) == self.cap:
            os, od, ot = self.q.popleft()
            self.present.discard((os, od, ot))
            self._advance_dest_front(od)

        # Insert new packet
        self.q.append((source, destination, timestamp))
        self.present.add(key)
        self._ensure_dest(destination)
        self.per_dest[destination]["ts"].append(timestamp)
        return True

    def forwardPacket(self) -> List[int]:
        if not self.q:
            return []
        s, d, t = self.q.popleft()
        self.present.discard((s, d, t))
        self._advance_dest_front(d)
        return [s, d, t]

    def getCount(self, destination: int, startTime: int, endTime: int) -> int:
        if destination not in self.per_dest:
            return 0
        b = self.per_dest[destination]
        ts, st = b["ts"], b["start"]
        # search only in the valid suffix ts[st:]
        L = bisect_left(ts, startTime, lo=st)
        R = bisect_right(ts, endTime, lo=st)
        return max(0, R - L)
