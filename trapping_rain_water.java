class Solution {
    public int trap(int[] height) {
        if(height == null || height.length <= 2) {
            return 0;
        }
        
        int left = 0, right = height.length - 1;
        int leftMax = 0, rightMax = 0;
        int trappedWater = 0;
        
        while(left < right) {
            leftMax = Math.max(leftMax, height[left]);
            rightMax = Math.max(rightMax, height[right]);
            
            // Water trapped at any point is the difference between the current height and min(leftMax, rightMax)
            if(leftMax < rightMax) {
                trappedWater += leftMax - height[left];
                left++;
            } else {
                trappedWater += rightMax - height[right];
                right--;
            }
        }
        return trappedWater;
    }
}
