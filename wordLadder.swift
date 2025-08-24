class Solution {
    func findLadders(_ beginWord: String, _ endWord: String, _ wordList: [String]) -> [[String]] {
        var wordSet = Set(wordList)
        guard wordSet.contains(endWord) else { return [] }

        func neighbors(_ s: String) -> [String] {
            var res: [String] = []
            var chars = Array(s)
            for i in 0..<chars.count {
                let old = chars[i]
                for a in 97...122 {
                    let ch = Character(UnicodeScalar(a)!)
                    if ch == old { continue }
                    chars[i] = ch
                    let cand = String(chars)
                    if wordSet.contains(cand) { res.append(cand) }
                }
                chars[i] = old
            }
            return res
        }

        var parents: [String: [String]] = [:]
        var level: Set<String> = [beginWord]
        var visited: Set<String> = [beginWord]
        var found = false

        while !level.isEmpty && !found {
            var nextLevel = Set<String>()
            for w in level {
                for v in neighbors(w) where !visited.contains(v) {
                    parents[v, default: []].append(w)
                    if !nextLevel.contains(v) { nextLevel.insert(v) }
                    if v == endWord { found = true }
                }
            }
            visited.formUnion(nextLevel)
            level = nextLevel
        }

        guard parents[endWord] != nil else { return [] }

        var results: [[String]] = []
        var path: [String] = [endWord]

        func dfs(_ word: String) {
            if word == beginWord {
                results.append(path.reversed())
                return
            }
            for p in parents[word] ?? [] {
                path.append(p)
                dfs(p)
                path.removeLast()
            }
        }

        dfs(endWord)
        return results
    }
}
