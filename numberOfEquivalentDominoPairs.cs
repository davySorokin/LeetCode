public class Solution {
    public int NumEquivDominoPairs(int[][] dominoes) {
        var map = new Dictionary<int, int>();
        int count = 0;

        foreach (var d in dominoes) {
            int a = Math.Min(d[0], d[1]);
            int b = Math.Max(d[0], d[1]);
            int key = a * 10 + b;

            if (map.ContainsKey(key)) {
                count += map[key];
                map[key]++;
            } else {
                map[key] = 1;
            }
        }

        return count;
    }
}
