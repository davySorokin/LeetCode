class Solution {
    int postIndex;
    Map<Integer, Integer> inorderIndexMap;

    public TreeNode buildTree(int[] inorder, int[] postorder) {
        postIndex = postorder.length - 1;
        inorderIndexMap = new HashMap<>();
        for (int i = 0; i < inorder.length; i++) {
            inorderIndexMap.put(inorder[i], i);
        }
        return buildTree(inorder, postorder, 0, inorder.length - 1);
    }

    private TreeNode buildTree(int[] inorder, int[] postorder, int left, int right) {
        if (left > right) return null;

        int rootValue = postorder[postIndex--];
        TreeNode root = new TreeNode(rootValue);

        // Build the right subtree before the left subtree
        root.right = buildTree(inorder, postorder, inorderIndexMap.get(rootValue) + 1, right);
        root.left = buildTree(inorder, postorder, left, inorderIndexMap.get(rootValue) - 1);

        return root;
    }
}
