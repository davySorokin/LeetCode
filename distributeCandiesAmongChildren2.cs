public class Solution {
    public long DistributeCandies(int n, int limit) {
        long count = 0;

        int startA = Math.Max(0, n - 2 * limit);
        int endA = Math.Min(limit, n);

        for (int a = startA; a <= endA; a++) {
            int s = n - a;
            int low = Math.Max(0, s - limit);
            int high = Math.Min(limit, s);
            count += high - low + 1;
        }

        return count;
    }
}
