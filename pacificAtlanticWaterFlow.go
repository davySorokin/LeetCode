package main

func pacificAtlantic(heights [][]int) [][]int {
    m, n := len(heights), len(heights[0])
    pacific := make([][]bool, m)
    atlantic := make([][]bool, m)
    for i := 0; i < m; i++ {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }

    dirs := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}
    var dfs func(int, int, [][]bool, int)
    dfs = func(r, c int, visited [][]bool, prev int) {
        if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] || heights[r][c] < prev {
            return
        }
        visited[r][c] = true
        for _, d := range dirs {
            dfs(r+d[0], c+d[1], visited, heights[r][c])
        }
    }

    for i := 0; i < m; i++ {
        dfs(i, 0, pacific, heights[i][0])
        dfs(i, n-1, atlantic, heights[i][n-1])
    }
    for j := 0; j < n; j++ {
        dfs(0, j, pacific, heights[0][j])
        dfs(m-1, j, atlantic, heights[m-1][j])
    }

    res := [][]int{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pacific[i][j] && atlantic[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }
    return res
}
