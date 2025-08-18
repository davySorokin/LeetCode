/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    var res [][]int
    var path []int

    var dfs func(node *TreeNode, remain int)
    dfs = func(node *TreeNode, remain int) {
        if node == nil {
            return
        }
        path = append(path, node.Val)
        remain -= node.Val

        if node.Left == nil && node.Right == nil && remain == 0 {
            cp := make([]int, len(path))
            copy(cp, path)
            res = append(res, cp)
        }

        dfs(node.Left, remain)
        dfs(node.Right, remain)

        path = path[:len(path)-1] // backtrack
    }

    dfs(root, targetSum)
    return res
}
