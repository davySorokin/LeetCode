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
    public void flatten(TreeNode root) {
        if (root == null) return;
        
        // Flatten the left and right subtree
        flatten(root.left);
        flatten(root.right);
        
        // Store the right subtree
        TreeNode rightSubtree = root.right;
        
        // Replace the right subtree with the left one
        root.right = root.left;
        root.left = null; // set the left child to null
        
        // Find the rightmost node
        TreeNode current = root;
        while (current.right != null) current = current.right;
        
        // Connect the stored right subtree
        current.right = rightSubtree;
    }
}
