func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    stack1 := []*TreeNode{}
    stack2 := []*TreeNode{}
    stack1 = append(stack1, root)

    // Traverse the tree
    for len(stack1) > 0 {
        node := stack1[len(stack1)-1]
        stack1 = stack1[:len(stack1)-1]
        stack2 = append(stack2, node)

        // Push left and right children to stack1
        if node.Left != nil {
            stack1 = append(stack1, node.Left)
        }
        if node.Right != nil {
            stack1 = append(stack1, node.Right)
        }
    }

    // Retrieve results from stack2 in postorder
    result := make([]int, len(stack2))
    for i := len(stack2) - 1; i >= 0; i-- {
        result[len(stack2)-1-i] = stack2[i].Val
    }

    return result
}
