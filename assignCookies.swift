class Solution {
    func findContentChildren(_ g: [Int], _ s: [Int]) -> Int {
        let greed  = g.sorted()
        let sizes  = s.sorted()
        var i = 0, j = 0, happy = 0
        
        while i < greed.count && j < sizes.count {
            if sizes[j] >= greed[i] {
                happy += 1
                i += 1
                j += 1
            } else {
                j += 1
            }
        }
        return happy
    }
}
