func snakesAndLadders(board [][]int) int {
    n := len(board)
    flatten := make([]int, n*n+1)
    idx := 1
    leftToRight := true
    for i := n - 1; i >= 0; i-- {
        row := board[i]
        if !leftToRight {
            reverse(row)
        }
        for _, val := range row {
            flatten[idx] = val
            idx++
        }
        leftToRight = !leftToRight
    }

    visited := make([]bool, n*n+1)
    queue := [][]int{{1, 0}} // {square, moves}
    visited[1] = true

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        for i := 1; i <= 6; i++ {
            next := curr[0] + i
            if next > n*n {
                break
            }
            if flatten[next] != -1 {
                next = flatten[next]
            }
            if next == n*n {
                return curr[1] + 1
            }
            if !visited[next] {
                visited[next] = true
                queue = append(queue, []int{next, curr[1] + 1})
            }
        }
    }
    return -1
}

func reverse(row []int) {
    for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
        row[i], row[j] = row[j], row[i]
    }
}
