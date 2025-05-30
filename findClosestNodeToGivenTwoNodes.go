func closestMeetingNode(edges []int, node1 int, node2 int) int {
    getDistances := func(start int) map[int]int {
        dist := make(map[int]int)
        visited := make(map[int]bool)
        steps := 0
        for start != -1 && !visited[start] {
            dist[start] = steps
            visited[start] = true
            start = edges[start]
            steps++
        }
        return dist
    }

    dist1 := getDistances(node1)
    dist2 := getDistances(node2)

    minDist := int(1e9)
    res := -1

    for node, d1 := range dist1 {
        if d2, ok := dist2[node]; ok {
            maxDist := max(d1, d2)
            if maxDist < minDist || (maxDist == minDist && node < res) {
                minDist = maxDist
                res = node
            }
        }
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
