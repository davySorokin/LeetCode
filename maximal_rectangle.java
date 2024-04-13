class Solution {
    public int maximalRectangle(char[][] matrix) {
        if (matrix.length == 0) return 0;
        int maxArea = 0;
        int[] height = new int[matrix[0].length];

        for (char[] row : matrix) {
            // Update the height of the histogram
            for (int i = 0; i < matrix[0].length; i++) {
                height[i] = row[i] == '1' ? height[i] + 1 : 0;
            }
            // Compute the largest rectangle from the histogram
            maxArea = Math.max(maxArea, largestRectangleArea(height));
        }

        return maxArea;
    }

    private int largestRectangleArea(int[] heights) {
        Stack<Integer> stack = new Stack<>();
        int maxArea = 0;
        int i = 0;
        while (i <= heights.length) {
            int h = (i == heights.length) ? 0 : heights[i];
            if (stack.isEmpty() || h >= heights[stack.peek()]) {
                stack.push(i++);
            } else {
                int height = heights[stack.pop()];
                int width = stack.isEmpty() ? i : i - 1 - stack.peek();
                maxArea = Math.max(maxArea, height * width);
            }
        }
        return maxArea;
    }
}
