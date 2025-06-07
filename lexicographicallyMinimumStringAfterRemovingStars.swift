class Solution {
    func clearStars(_ s: String) -> String {
        let chars = Array(s)
        let n = chars.count
        var g = Array(repeating: [Int](), count: 26)
        var deleted = Array(repeating: false, count: n)

        for i in 0..<n {
            if chars[i] == "*" {
                deleted[i] = true
                for j in 0..<26 {
                    if !g[j].isEmpty {
                        let idx = g[j].removeLast()
                        deleted[idx] = true
                        break
                    }
                }
            } else {
                let idx = Int(chars[i].asciiValue! - Character("a").asciiValue!)
                g[idx].append(i)
            }
        }

        var result = ""
        for i in 0..<n {
            if !deleted[i] && chars[i] != "*" {
                result.append(chars[i])
            }
        }

        return result
    }
}
