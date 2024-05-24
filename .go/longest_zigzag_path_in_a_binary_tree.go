package main

import (
	"fmt"
)

// Assume TreeNode is already defined in the imported package

func longestZigZag(root *TreeNode) int {
    var dfs func(node *TreeNode, direction bool, length int)
    maxLen := 0
    
    dfs = func(node *TreeNode, direction bool, length int) {
        if node == nil {
            return
        }
        maxLen = max(maxLen, length)
        if direction {
            dfs(node.Left, !direction, length+1)
            dfs(node.Right, direction, 1)
        } else {
            dfs(node.Right, !direction, length+1)
            dfs(node.Left, direction, 1)
        }
    }
    
    dfs(root, true, 0)  // Start by moving right
    dfs(root, false, 0) // Start by moving left
    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
