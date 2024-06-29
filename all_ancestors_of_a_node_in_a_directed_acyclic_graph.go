package main

import (
	"fmt"
	"sort"
)

func getAncestors(n int, edges [][]int) [][]int {
	// Initialize the graph and indegree array
	graph := make(map[int][]int)
	indegree := make([]int, n)
	ancestors := make([]map[int]struct{}, n)

	for i := 0; i < n; i++ {
		ancestors[i] = make(map[int]struct{})
	}

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	// Topological sort using Kahn's algorithm
	queue := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[node] {
			// Add current node's ancestors to the neighbor's ancestors
			for ancestor := range ancestors[node] {
				ancestors[neighbor][ancestor] = struct{}{}
			}
			ancestors[neighbor][node] = struct{}{}

			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Convert ancestor sets to sorted slices
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		for ancestor := range ancestors[i] {
			result[i] = append(result[i], ancestor)
		}
		sort.Ints(result[i])
	}

	return result
}
