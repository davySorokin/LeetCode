/*type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}*/ 
// <--- Struct NOT needed in Leetcode (already defined)

func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return buildTrees(1, n)
}

func buildTrees(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }

    var allTrees []*TreeNode
    for i := start; i <= end; i++ {
        // Generate all possible left and right subtrees
        leftTrees := buildTrees(start, i-1)
        rightTrees := buildTrees(i+1, end)

        // Connect left and right subtrees to the root i
        for _, left := range leftTrees {
            for _, right := range rightTrees {
                root := &TreeNode{Val: i}
                root.Left = left
                root.Right = right
                allTrees = append(allTrees, root)
            }
        }
    }
    return allTrees
}
