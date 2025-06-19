package main

import "fmt"

func generateParenthesis(n int) []string {
    var res []string
    var dfs func(open, close int, cur []byte)
    dfs = func(open, close int, cur []byte) {
        if len(cur) == 2*n {
            res = append(res, string(cur))
            return
        }
        if open < n {
            dfs(open+1, close, append(cur, '('))
        }
        if close < open {
            dfs(open, close+1, append(cur, ')'))
        }
    }
    dfs(0, 0, make([]byte, 0, 2*n))
    return res
}

func main() {
    fmt.Println(generateParenthesis(3))
}
