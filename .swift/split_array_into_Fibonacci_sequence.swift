class Solution {
    func splitIntoFibonacci(_ num: String) -> [Int] {
        var result = [Int]()
        backtrack(num, index: 0, &result)
        return result
    }

    private func backtrack(_ num: String, index: Int, _ path: inout [Int]) -> Bool {
        if index == num.count && path.count >= 3 {
            return true
        }

        var number: Int64 = 0
        let startIndex = num.index(num.startIndex, offsetBy: index)
        for i in index..<num.count {
            let endIndex = num.index(num.startIndex, offsetBy: i)
            let char = num[endIndex]
            if num[startIndex] == "0" && i > index { break }  // Prevent "leading zero" numbers except "0" itself.
            number = number * 10 + Int64(String(char))!
            if number > Int32.max { break }  // Avoid overflow

            if path.count >= 2 && Int(number) != path[path.count - 1] + path[path.count - 2] {
                continue
            }

            if path.count <= 1 || Int(number) == path[path.count - 1] + path[path.count - 2] {
                path.append(Int(number))
                if backtrack(num, index: i + 1, &path) {
                    return true
                }
                path.removeLast()
            }
        }
        return false
    }
}
