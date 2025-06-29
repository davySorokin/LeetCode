class Solution {
    func convert(_ s: String, _ numRows: Int) -> String {
        if numRows == 1 || numRows >= s.count {
            return s
        }

        var rows = Array(repeating: "", count: numRows)
        var currRow = 0
        var goingDown = false

        for c in s {
            rows[currRow].append(c)
            if currRow == 0 || currRow == numRows - 1 {
                goingDown.toggle()
            }
            currRow += goingDown ? 1 : -1
        }

        return rows.joined()
    }
}
