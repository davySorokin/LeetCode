/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    // simple slice-backed queue
    q := []*TreeNode{root}
    res := [][]int{}

    for len(q) > 0 {
        size := len(q)
        level := make([]int, 0, size)
        for i := 0; i < size; i++ {
            node := q[0]
            q = q[1:]
            level = append(level, node.Val)
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        res = append(res, level)
    }

    // reverse to get bottom-up
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
    return res
}
