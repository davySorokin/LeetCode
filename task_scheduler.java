import java.util.Arrays;

class Solution {
    public int leastInterval(char[] tasks, int n) {
        if (n == 0) {
            return tasks.length;
        }
        
        // Frequency of tasks
        int[] frequencies = new int[26];
        for (char task : tasks) {
            frequencies[task - 'A']++;
        }
        
        Arrays.sort(frequencies);
        
        // Max frequency
        int maxFreq = frequencies[25];
        int idleTime = (maxFreq - 1) * n;
        
        for (int i = 24; i >= 0 && frequencies[i] > 0 && idleTime > 0; --i) {
            idleTime -= Math.min(maxFreq - 1, frequencies[i]);
        }
        idleTime = Math.max(0, idleTime);
        
        return tasks.length + idleTime;
    }
}
