class Solution {
    public int averageOfSubtree(TreeNode root) {
        return countNodesWithAverage(root).nodeCount;
    }

    private Result countNodesWithAverage(TreeNode node) {
        if (node == null) {
            return new Result(0, 0, 0);
        }

        Result leftResult = countNodesWithAverage(node.left);
        Result rightResult = countNodesWithAverage(node.right);

        int sum = node.val + leftResult.sum + rightResult.sum;
        int count = 1 + leftResult.count + rightResult.count;
        
        // Check if current node's value is equal to the average of its subtree
        int nodeCount = (node.val == sum / count) ? 1 : 0;

        // Total count is sum of the counts from left and right subtrees plus the current node if it matches
        nodeCount += leftResult.nodeCount + rightResult.nodeCount;

        return new Result(sum, count, nodeCount);
    }

    private static class Result {
        int sum; // Sum of all nodes in the subtree
        int count; // Count of all nodes in the subtree
        int nodeCount; // Count of nodes that match the average condition

        Result(int sum, int count, int nodeCount) {
            this.sum = sum;
            this.count = count;
            this.nodeCount = nodeCount;
        }
    }
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}
