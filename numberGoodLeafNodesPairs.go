/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countPairs(root *TreeNode, distance int) int {
    var result int
    
    // Helper function that returns an array representing distances to leaf nodes
    var dfs func(node *TreeNode) []int
    dfs = func(node *TreeNode) []int {
        if node == nil {
            return []int{}
        }
        
        // If it's a leaf node, return an array with a single distance 1
        if node.Left == nil && node.Right == nil {
            return []int{1}
        }
        
        // Get distances from left and right children
        leftDistances := dfs(node.Left)
        rightDistances := dfs(node.Right)
        
        // Check all pairs from left and right distances
        for _, ld := range leftDistances {
            for _, rd := range rightDistances {
                if ld + rd <= distance {
                    result++
                }
            }
        }
        
        // Increment distances by 1 and filter by distance
        var newDistances []int
        for _, d := range leftDistances {
            if d+1 < distance {
                newDistances = append(newDistances, d+1)
            }
        }
        for _, d := range rightDistances {
            if d+1 < distance {
                newDistances = append(newDistances, d+1)
            }
        }
        
        return newDistances
    }
    
    dfs(root)
    return result
}
