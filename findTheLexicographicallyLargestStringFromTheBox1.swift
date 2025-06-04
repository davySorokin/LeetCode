class Solution {
    func answerString(_ word: String, _ numFriends: Int) -> String {
        var result = [String]()
        let chars = Array(word)
        
        func backtrack(_ index: Int, _ splitsLeft: Int, _ path: [String]) {
            if splitsLeft == 1 {
                let remaining = String(chars[index...])
                result.append(contentsOf: path + [remaining])
                return
            }
            
            for i in index..<chars.count - splitsLeft + 1 {
                let part = String(chars[index...i])
                backtrack(i + 1, splitsLeft - 1, path + [part])
            }
        }
        
        backtrack(0, numFriends, [])
        return result.max() ?? ""
    }
}
