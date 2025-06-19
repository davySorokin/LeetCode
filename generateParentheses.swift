class Solution {
    func generateParenthesis(_ n: Int) -> [String] {
        var result = [String]()
        var path = [Character]()
        
        func dfs(_ open: Int, _ close: Int) {
            if path.count == 2 * n {
                result.append(String(path))
                return
            }
            if open < n {
                path.append("(")
                dfs(open + 1, close)
                path.removeLast()
            }
            if close < open {
                path.append(")")
                dfs(open, close + 1)
                path.removeLast()
            }
        }
        
        dfs(0, 0)
        return result
    }
}
