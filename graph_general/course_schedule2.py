from typing import List
from collections import deque, defaultdict

class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        # Create a graph in adjacency list form and a list to count prerequisites
        graph = defaultdict(list)
        indegree = [0] * numCourses
        
        # Construct the graph and indegree array
        for dest, src in prerequisites:
            graph[src].append(dest)
            indegree[dest] += 1
        
        # Queue for BFS
        queue = deque([i for i in range(numCourses) if indegree[i] == 0])
        order = []
        
        # Process nodes with no prerequisites
        while queue:
            node = queue.popleft()
            order.append(node)
            for neighbor in graph[node]:
                indegree[neighbor] -= 1
                if indegree[neighbor] == 0:
                    queue.append(neighbor)
        
        return order if len(order) == numCourses else []
