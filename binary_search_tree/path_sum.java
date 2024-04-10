/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public boolean hasPathSum(TreeNode root, int targetSum) {
        // Base case: when we reach a null node
        if (root == null) {
            return false;
        }
        
        // Check if we're at a leaf node and if the path sum equals the target sum
        if (root.left == null && root.right == null && root.val == targetSum) {
            return true;
        }
        
        // Recursively check the left and right subtree with the updated sum
        return hasPathSum(root.left, targetSum - root.val) || 
               hasPathSum(root.right, targetSum - root.val);
    }
}
