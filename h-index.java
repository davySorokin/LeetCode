class Solution {
    public int hIndex(int[] citations) {
        Arrays.sort(citations);
        int n = citations.length;
        int h = 0;
        for (int i = 0; i < n; i++) {
            int currH = Math.min(citations[i], n - i);
            if (currH > h) {
                h = currH;
            }
        }
        return h;
    }
}
