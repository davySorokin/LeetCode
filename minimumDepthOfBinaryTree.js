/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val);
 *     this.left = (left===undefined ? null : left);
 *     this.right = (right===undefined ? null : right);
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
var minDepth = function(root) {
    if (root === null) return 0; // If the root is null, return depth as 0

    // If one of the subtrees is null, only consider the non-null subtree
    if (root.left === null) {
        return 1 + minDepth(root.right);
    }
    if (root.right === null) {
        return 1 + minDepth(root.left);
    }

    // If both subtrees are non-null, return the minimum depth between them
    return 1 + Math.min(minDepth(root.left), minDepth(root.right));
};
