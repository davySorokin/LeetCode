package main

import (
	"fmt"
)

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// Create an adjacency list to represent the graph
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Use BFS to find if there is a path from source to destination
	visited := make([]bool, n)
	queue := []int{source}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == destination {
			return true
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}
