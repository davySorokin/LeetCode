class Solution {
    public int removeElement(int[] nums, int val) {
        int insertPosition = 0;
        for (int num : nums) {
            if (num != val) {
                nums[insertPosition] = num;
                insertPosition++;
            }
        }
        return insertPosition;
    }
}
