func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
    n := len(edges1) + 1
    m := len(edges2) + 1

    tree1 := make([][]int, n)
    for _, e := range edges1 {
        a, b := e[0], e[1]
        tree1[a] = append(tree1[a], b)
        tree1[b] = append(tree1[b], a)
    }

    tree2 := make([][]int, m)
    for _, e := range edges2 {
        a, b := e[0], e[1]
        tree2[a] = append(tree2[a], b)
        tree2[b] = append(tree2[b], a)
    }

    computeEvenOddCounts := func(tree [][]int) (even []int, odd []int) {
        size := len(tree)
        even = make([]int, size)
        odd = make([]int, size)
        visited := make([]bool, size)

        var dfs func(node, depth int)
        dfs = func(node, depth int) {
            visited[node] = true
            if depth%2 == 0 {
                even[0]++
            } else {
                odd[0]++
            }
            for _, nei := range tree[node] {
                if !visited[nei] {
                    dfs(nei, depth+1)
                }
            }
        }

        dfs(0, 0)

        evenCount := make([]int, size)
        oddCount := make([]int, size)

        visited = make([]bool, size)
        var dfsCount func(node, depth int)
        dfsCount = func(node, depth int) {
            visited[node] = true
            if depth%2 == 0 {
                evenCount[node] = even[0]
                oddCount[node] = odd[0]
            } else {
                evenCount[node] = odd[0]
                oddCount[node] = even[0]
            }
            for _, nei := range tree[node] {
                if !visited[nei] {
                    dfsCount(nei, depth+1)
                }
            }
        }

        dfsCount(0, 0)
        return evenCount, oddCount
    }

    even1, _ := computeEvenOddCounts(tree1)
    _, odd2 := computeEvenOddCounts(tree2)

    maxOdd2 := 0
    for _, v := range odd2 {
        if v > maxOdd2 {
            maxOdd2 = v
        }
    }

    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = even1[i] + maxOdd2
    }

    return res
}
