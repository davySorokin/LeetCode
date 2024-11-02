func countSubTrees(n int, edges [][]int, labels string) []int {
	// Build adjacency list
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := make([]int, n)
	visited := make([]bool, n)

	// DFS function to traverse tree and count labels
	var dfs func(node int) []int
	dfs = func(node int) []int {
		count := make([]int, 26)
		visited[node] = true
		labelIdx := labels[node] - 'a'
		count[labelIdx]++

		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				subCount := dfs(neighbor)
				for i := 0; i < 26; i++ {
					count[i] += subCount[i]
				}
			}
		}
		ans[node] = count[labelIdx]
		return count
	}

	// Start DFS from root node (0)
	dfs(0)

	return ans
}
