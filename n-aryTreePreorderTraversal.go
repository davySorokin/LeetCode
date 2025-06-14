/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
        return nil
    }
    out := []int{root.Val}
    for _, c := range root.Children {
        out = append(out, preorder(c)...)
    }
    return out
}
