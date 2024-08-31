package main

import (
	"fmt"
)

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	if n == 1 {
		return 0
	}

	// Initialize the queue and the visited map
	type state struct {
		node int
		mask int
		dist int
	}
	queue := []state{}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, 1<<n)
	}

	// Start BFS from every node
	for i := 0; i < n; i++ {
		mask := 1 << i
		queue = append(queue, state{i, mask, 0})
		visited[i][mask] = true
	}

	// Perform BFS
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Check if all nodes are visited
		if curr.mask == (1<<n)-1 {
			return curr.dist
		}

		// Visit all the neighbors
		for _, neighbor := range graph[curr.node] {
			nextMask := curr.mask | (1 << neighbor)
			if !visited[neighbor][nextMask] {
				visited[neighbor][nextMask] = true
				queue = append(queue, state{neighbor, nextMask, curr.dist + 1})
			}
		}
	}

	return -1
}
