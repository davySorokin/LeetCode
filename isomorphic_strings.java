class Solution {
    public boolean isIsomorphic(String s, String t) {
        // Create two arrays to store the last seen positions of characters.
        int[] mappingST = new int[256]; // For mapping of characters of s to t
        int[] mappingTS = new int[256]; // For mapping of characters of t to s

        // Initialise the arrays with -1, indicating that no mappings have been recorded yet.
        for (int i = 0; i < 256; i++) {
            mappingST[i] = -1;
            mappingTS[i] = -1;
        }

        for (int i = 0; i < s.length(); i++) {
            char c1 = s.charAt(i);
            char c2 = t.charAt(i);

            // Check if the current mapping is consistent with previous mappings.
            if (mappingST[c1] != -1 && mappingST[c1] != c2 || mappingTS[c2] != -1 && mappingTS[c2] != c1) {
                return false;
            }

            // Record the new mapping.
            mappingST[c1] = c2;
            mappingTS[c2] = c1;
        }

        return true;
    }
}
