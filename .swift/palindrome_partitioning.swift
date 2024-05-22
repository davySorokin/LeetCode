class Solution {
    func partition(_ s: String) -> [[String]] {
        var result = [[String]]()
        var currentPartition = [String]()
        backtrack(0, &currentPartition, &result, s)
        return result
    }

    private func backtrack(_ start: Int, _ currentPartition: inout [String], _ result: inout [[String]], _ s: String) {
        if start >= s.count { // If reached the end of the string
            result.append(currentPartition)
            return
        }

        for end in start..<s.count {
            let startIdx = s.index(s.startIndex, offsetBy: start)
            let endIdx = s.index(s.startIndex, offsetBy: end)
            let substring = String(s[startIdx...endIdx])

            if isPalindrome(substring) {
                currentPartition.append(substring)
                backtrack(end + 1, &currentPartition, &result, s)
                currentPartition.removeLast()
            }
        }
    }

    private func isPalindrome(_ s: String) -> Bool {
        return s == String(s.reversed())
    }
}
