from typing import List

class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        def dfs(course):
            if visited[course] == -1:  # cycle detected
                return False
            if visited[course] == 1:  # already visited and no cycle
                return True
            visited[course] = -1  # mark as being visited
            for prereq in adj_list[course]:
                if not dfs(prereq):
                    return False
            visited[course] = 1  # mark as visit completed
            return True
        
        # Create an adjacency list for the graph representation
        adj_list = [[] for _ in range(numCourses)]
        for dest, src in prerequisites:
            adj_list[src].append(dest)
        
        # List to track visit status: 0 = not visited, -1 = being visited, 1 = visited
        visited = [0] * numCourses
        for course in range(numCourses):
            if visited[course] == 0:
                if not dfs(course):
                    return False  # if cycle is detected, return False
        return True  # if no cycle is detected in any DFS, return True
