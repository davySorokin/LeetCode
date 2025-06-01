public class Solution {
    public long DistributeCandies(int n, int limit) {
        long count = 0;

        for (int a = 0; a <= limit; a++) {
            for (int b = 0; b <= limit; b++) {
                int c = n - a - b;
                if (c >= 0 && c <= limit) {
                    count++;
                }
            }
        }

        return count;
    }
}
