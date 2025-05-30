public class Solution {
    public int ClosestMeetingNode(int[] edges, int node1, int node2) {
        int n = edges.Length;

        int[] GetDistances(int start) {
            int[] dist = new int[n];
            Array.Fill(dist, -1);
            int steps = 0;
            while (start != -1 && dist[start] == -1) {
                dist[start] = steps++;
                start = edges[start];
            }
            return dist;
        }

        int[] dist1 = GetDistances(node1);
        int[] dist2 = GetDistances(node2);

        int minDist = int.MaxValue;
        int result = -1;

        for (int i = 0; i < n; i++) {
            if (dist1[i] != -1 && dist2[i] != -1) {
                int maxDist = Math.Max(dist1[i], dist2[i]);
                if (maxDist < minDist || (maxDist == minDist && i < result)) {
                    minDist = maxDist;
                    result = i;
                }
            }
        }

        return result;
    }
}
