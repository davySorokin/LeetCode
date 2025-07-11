/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public var val: Int
 *     public var left: TreeNode?
 *     public var right: TreeNode?
 *     public init() { self.val = 0; self.left = nil; self.right = nil; }
 *     public init(_ val: Int) { self.val = val; self.left = nil; self.right = nil; }
 *     public init(_ val: Int, _ left: TreeNode?, _ right: TreeNode?) {
 *         self.val = val
 *         self.left = left
 *         self.right = right
 *     }
 * }
 */
class Solution {
    func recoverTree(_ root: TreeNode?) {
        var first: TreeNode?
        var second: TreeNode?
        var prev: TreeNode?

        func inorder(_ node: TreeNode?) {
            guard let node = node else { return }

            inorder(node.left)

            if let p = prev, p.val > node.val {
                if first == nil { first = p }
                second = node
            }
            prev = node

            inorder(node.right)
        }

        inorder(root)
        if let f = first, let s = second {
            (f.val, s.val) = (s.val, f.val)
        }
    }
}
