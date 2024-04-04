class Solution {
    public int maxDepth(String s) {
        int currentDepth = 0;
        int maxDepth = 0;

        for (int i = 0; i < s.length(); i++) {
            char ch = s.charAt(i);
            if (ch == '(') {
                currentDepth++;  // increment the depth for each open parenthesis
                maxDepth = Math.max(maxDepth, currentDepth);  // update max depth if needed
            } else if (ch == ')') {
                currentDepth--;  // decrement the depth for each closed parenthesis
            }
        }

        return maxDepth;
    }
}
