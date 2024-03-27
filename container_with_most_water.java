class Solution {
    public int maxArea(int[] height) {
        int maxArea = 0;
        int left = 0, right = height.length - 1;
        
        while (left < right) {
            // Calculate the area with the current boundary lines
            int width = right - left;
            int currentHeight = Math.min(height[left], height[right]);
            maxArea = Math.max(maxArea, width * currentHeight);
            
            // Move the pointer that points to the lower line
            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }
        }
        
        return maxArea;
    }
}
