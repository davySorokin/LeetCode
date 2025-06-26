class Solution {
    func imageSmoother(_ img: [[Int]]) -> [[Int]] {
        let m = img.count, n = img[0].count
        var res = Array(repeating: Array(repeating: 0, count: n), count: m)
        
        for i in 0..<m {
            for j in 0..<n {
                var sum = 0, count = 0
                for dx in -1...1 {
                    for dy in -1...1 {
                        let ni = i + dx, nj = j + dy
                        if ni >= 0 && ni < m && nj >= 0 && nj < n {
                            sum += img[ni][nj]
                            count += 1
                        }
                    }
                }
                res[i][j] = sum / count
            }
        }
        
        return res
    }
}
