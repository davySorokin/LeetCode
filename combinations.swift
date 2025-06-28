class Solution {
    func combine(_ n: Int, _ k: Int) -> [[Int]] {
        var result = [[Int]]()
        var path = [Int]()

        func backtrack(_ start: Int) {
            // If the current combination is complete
            if path.count == k {
                result.append(path)
                return
            }

            // If start > n, there's nothing to add
            if start > n {
                return
            }

            // Use i...n only if start <= n
            for i in start...n {
                path.append(i)
                backtrack(i + 1)
                path.removeLast()
            }
        }

        backtrack(1)
        return result
    }
}
