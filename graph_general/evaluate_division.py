from typing import List

class UnionFind:
    def __init__(self):
        self.parent = {}
        self.ratio = {}  # ratio of node / parent

    def find(self, x):
        if x not in self.parent:
            self.parent[x] = x
            self.ratio[x] = 1.0
        if x != self.parent[x]:
            origin = self.parent[x]
            self.parent[x], root_ratio = self.find(self.parent[x])
            self.ratio[x] *= root_ratio  # adjust the ratio to the root
        return self.parent[x], self.ratio[x]

    def union(self, x, y, value):
        root_x, ratio_x = self.find(x)
        root_y, ratio_y = self.find(y)
        if root_x != root_y:
            self.parent[root_x] = root_y
            self.ratio[root_x] = ratio_y * value / ratio_x

    def connected(self, x, y):
        if x in self.parent and y in self.parent:
            root_x, ratio_x = self.find(x)
            root_y, ratio_y = self.find(y)
            if root_x == root_y:
                return ratio_x / ratio_y
        return -1.0

class Solution:
    def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:
        uf = UnionFind()
        for (x, y), value in zip(equations, values):
            uf.union(x, y, value)
        
        return [uf.connected(q[0], q[1]) for q in queries]
