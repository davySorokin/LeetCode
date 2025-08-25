class Solution {
    func findDiagonalOrder(_ mat: [[Int]]) -> [Int] {
        let m = mat.count
        let n = mat[0].count
        var result: [Int] = []

        for d in 0..<(m + n - 1) {
            var intermediate: [Int] = []

            var r = d < n ? 0 : d - n + 1
            var c = d < n ? d : n - 1

            while r < m && c >= 0 {
                intermediate.append(mat[r][c])
                r += 1
                c -= 1
            }

            if d % 2 == 0 {
                result.append(contentsOf: intermediate.reversed())
            } else {
                result.append(contentsOf: intermediate)
            }
        }

        return result
    }
}
