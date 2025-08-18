/**
 * Definition for a binary tree node.
 * class TreeNode {
 *   int val;
 *   TreeNode? left;
 *   TreeNode? right;
 *   TreeNode([this.val = 0, this.left, this.right]);
 * }
 */
class Solution {
  List<List<int>> pathSum(TreeNode? root, int targetSum) {
    final res = <List<int>>[];
    final path = <int>[];

    void dfs(TreeNode? node, int remain) {
      if (node == null) return;

      path.add(node.val);
      remain -= node.val;

      final isLeaf = node.left == null && node.right == null;
      if (isLeaf && remain == 0) {
        res.add(List<int>.from(path)); // record a copy
      }

      dfs(node.left, remain);
      dfs(node.right, remain);

      path.removeLast(); // backtrack
    }

    dfs(root, targetSum);
    return res;
  }
}
