class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        // Indexes for the last elements of nums1, nums2, and the merged array
        int i = m - 1, j = n - 1, k = m + n - 1;

        // Merge nums2 into nums1 starting from the end of each array
        while (j >= 0) {
            // If nums1 is exhausted or the current element of nums2 is greater,
            // copy the element from nums2
            if (i < 0 || nums2[j] > nums1[i]) {
                nums1[k--] = nums2[j--];
            } else { // Otherwise, copy the element from nums1
                nums1[k--] = nums1[i--];
            }
        }
    }
}
