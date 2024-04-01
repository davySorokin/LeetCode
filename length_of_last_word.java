class Solution {
    public int lengthOfLastWord(String s) {
        int length = 0;
        boolean charFound = false;

        // Start from the end of the string
        for (int i = s.length() - 1; i >= 0; i--) {
            if (s.charAt(i) != ' ') {
                // If the current character is not a space, increment length
                length++;
                charFound = true;
            } else if (charFound) {
                // If we've found the end of the last word, exit the loop
                break;
            }
        }

        return length;
    }
}
