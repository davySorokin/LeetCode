class Solution {
    func matrixReshape(_ mat: [[Int]], _ r: Int, _ c: Int) -> [[Int]] {
        let m = mat.count
        let n = mat[0].count
        
        // Check if reshaping is possible
        if m * n != r * c {
            return mat
        }
        
        // Flatten the matrix into a single array
        var flattened = [Int]()
        for i in 0..<m {
            for j in 0..<n {
                flattened.append(mat[i][j])
            }
        }
        
        // Create the reshaped matrix
        var reshaped = [[Int]]()
        for i in 0..<r {
            var row = [Int]()
            for j in 0..<c {
                row.append(flattened[i * c + j])
            }
            reshaped.append(row)
        }
        
        return reshaped
    }
}
