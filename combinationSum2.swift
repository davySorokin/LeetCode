class Solution {
    func combinationSum2(_ candidates: [Int], _ target: Int) -> [[Int]] {
        let candidates = candidates.sorted()
        var result = [[Int]]()
        
        func backtrack(_ start: Int, _ target: Int, _ path: [Int]) {
            if target == 0 {
                result.append(path)
                return
            }
            
            for i in start..<candidates.count {
                if i > start && candidates[i] == candidates[i - 1] {
                    continue // skip duplicates
                }
                if candidates[i] > target {
                    break
                }
                backtrack(i + 1, target - candidates[i], path + [candidates[i]])
            }
        }
        
        backtrack(0, target, [])
        return result
    }
}
